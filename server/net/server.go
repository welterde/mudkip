package main

import "os"
import "net"
import "sync"
import "time"
import "log"
import "io"
import "fmt"
import "crypto/tls"
import "crypto/rand"
import "mudkip/lib"

const (
	// This should be the very first 6 bytes we receive on a new connection.
	// It allows us to filter out unrelated connections.
	ServerName = "MUDKIP"

	// This is the version by which a client can identify the server and see if
	// it is compatible. We send this to a client directly after we
	// receive the correct ServerName value. The version is updated everytime
	// this server changes in a way that will make it incompatible with older
	// versions.
	ServerVersion = byte(1)
)

type Server struct {
	config       *Config
	messages     chan lib.Message
	log          *log.Logger
	conn         net.Listener
	rwm          *sync.RWMutex
	clients      map[string]*Client
	clientclosed chan net.Addr
}

func NewServer(cfg *Config) *Server {
	s := new(Server)
	s.config = cfg

	var logtarget *os.File
	if cfg.LogFile != "" {
		var err os.Error
		if logtarget, err = os.Open(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0); err != nil {
			logtarget = os.Stdout
		}
	} else {
		logtarget = os.Stdout
	}

	s.log = log.New(logtarget, nil, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	s.rwm = new(sync.RWMutex)
	s.clients = make(map[string]*Client)
	s.messages = make(chan lib.Message, 32)
	s.clientclosed = make(chan net.Addr)
	return s
}

// Get the client associated with the given ID.
func (this *Server) GetClient(id string) *Client {
	if c, ok := this.clients[id]; ok {
		return c
	}
	return nil
}

func (this *Server) Messages() <-chan lib.Message { return this.messages }
func (this *Server) IsSecure() bool               { return this.config.Secure }

func (this *Server) Close() {
	this.Info("Shutting down...")

	if this.clients != nil {
		for id, _ := range this.clients {
			this.clients[id].Close()
		}
	}

	close(this.messages)

	if this.conn != nil {
		this.rwm.Lock()
		this.conn.Close()
		this.conn = nil
		this.rwm.Unlock()
		time.Sleep(1e9)
	}
}

func (this *Server) Open() (err os.Error) {
	if this.conn != nil {
		return
	}

	this.rwm.Lock()
	if this.conn, err = net.Listen("tcp", this.config.ListenAddr); err != nil {
		this.rwm.Unlock()
		return
	}
	this.rwm.Unlock()

	if this.config.Secure {
		cfg := new(tls.Config)
		cfg.Rand = rand.Reader
		cfg.Time = time.Nanoseconds

		cfg.Certificates = make([]tls.Certificate, 1)
		if cfg.Certificates[0], err = tls.LoadX509KeyPair(
			this.config.ServerCert,
			this.config.ServerKey,
		); err != nil {
			return
		}

		this.rwm.Lock()
		this.conn = tls.NewListener(this.conn, cfg)
		this.rwm.Unlock()
	}

	go this.poll()
	go this.clean()
	return
}

// Polls clientclosed channel for closed connections
func (this *Server) clean() {
	var addr net.Addr
	var id string

	for {
		select {
		case addr = <-this.clientclosed:
			if addr != nil {
				this.messages <- lib.NewClientDisconnected(addr)
				this.rwm.Lock()

				id = addr.String()
				if _, ok := this.clients[id]; ok {
					this.clients[id] = nil, false
				}

				this.rwm.Unlock()
			}
		}

		if closed(this.clientclosed) {
			return
		}
	}
}

func (this *Server) poll() {
	var err os.Error
	var client net.Conn

loop:
	for this.conn != nil {
		if client, err = this.conn.Accept(); err != nil {
			if this.conn == nil {
				return
			}

			this.Error("%T %v", err, err)
			continue loop
		}

		if len(this.clients) >= this.config.MaxClients {
			client.Write([]uint8{lib.MTMaxClientsReached})
			client.Close()
			continue loop
		}

		go this.process(client)
	}
}

func (this *Server) process(conn net.Conn) {
	var err os.Error
	var endpoint io.ReadWriteCloser

	if this.config.Secure {
		// FIXME: Half-arsed SSL implementation. Does not verify certificate.
		cf := new(tls.Config)
		cf.Rand = rand.Reader
		cf.Time = time.Nanoseconds
		endpoint = tls.Client(conn, cf)
	} else {
		endpoint = conn
	}

	// Send server version. Client can decide if it supports this server
	// version. If not, it should close the connection and go away.
	if _, err = endpoint.Write([]uint8{lib.MTServerVersion, ServerVersion}); err != nil {
		this.Error("%s %v", conn.RemoteAddr(), err)
		return
	}

	in := make([]byte, len(ServerName))

	// Make sure we have a relevant connection. Eg: it sends us the 'Magic
	// handshake'. if not, it's likely a connection that made a wrong turn
	// somewhere and doesn't belong here. Portscanners have a tendency to
	// get into this situation. They bombard us with a range of standard
	// service requests in the hopes of getting a useful response.
	if _, err = endpoint.Read(in); err != nil {
		this.Error("%s %v", conn.RemoteAddr(), err)
		return
	}

	if string(in) != ServerName {
		conn.Close()
		return
	}

	// Connection is relevant. We need to create a proper Client instance and
	// get rid of any clients from this source we already have. This can happen
	// when the client has disconnected/timed out for whatever reason and wants
	// to resume its session.
	id := conn.RemoteAddr().String()

	// Unlikely to happen, but you never know. If so, consider it a questionable
	// source and discard.
	if id == "" {
		conn.Close()
		return
	}

	//conn.SetKeepAlive(true)
	conn.SetTimeout(int64(this.config.ClientTimeout) * 6e10)

	// If client already exists, we have a reconnect. Close the old one.
	if _, ok := this.clients[id]; ok {
		this.clients[id].Close()
	}

	// Store new client.
	client := newClient(endpoint, conn.RemoteAddr(), this.clientclosed, this.messages)

	this.rwm.Lock()
	this.clients[id] = client
	this.rwm.Unlock()
	this.messages <- lib.NewClientConnected(conn.RemoteAddr())

	// Let's get this show on the road!
	go client.Run()
}

func (this *Server) Info(f string, a ...interface{}) {
	s := fmt.Sprintf(f, a...) // issue 1136
	this.rwm.Lock()
	this.log.Logf("[i] %s", s)
	this.rwm.Unlock()
}

func (this *Server) Error(f string, a ...interface{}) {
	s := fmt.Sprintf(f, a...) // issue 1136
	this.rwm.Lock()
	this.log.Logf("[e] %s", s)
	this.rwm.Unlock()
}

func (this *Server) Warn(f string, a ...interface{}) {
	s := fmt.Sprintf(f, a...) // issue 1136
	this.rwm.Lock()
	this.log.Logf("[w] %s", s)
	this.rwm.Unlock()
}

func (this *Server) Debug(f string, a ...interface{}) {
	s := fmt.Sprintf(f, a...) // issue 1136
	this.rwm.Lock()
	this.log.Logf("[d] %s", s)
	this.rwm.Unlock()
}
