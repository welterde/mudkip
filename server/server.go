package server

import "os"
import "net"
import "sync"
import "time"
import "log"
import "io"
import "fmt"
import "mudkip/lib"

const (
	ClientTimeout = 2 // minutes

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
	Messages     chan lib.Message
	log          *log.Logger
	secure       bool
	conn         *net.TCPListener
	rwm          *sync.RWMutex
	clients      map[string]*Client
	clientclosed chan net.Addr
	maxclients   int
}

func NewServer(logtarget io.Writer, secure bool, maxclients int) *Server {
	s := new(Server)
	s.log = log.New(logtarget, nil, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	s.rwm = new(sync.RWMutex)
	s.secure = secure
	s.clients = make(map[string]*Client)
	s.Messages = make(chan lib.Message, 32)
	s.clientclosed = make(chan net.Addr)
	s.maxclients = maxclients
	return s
}

// Get the client associated with the given ID.
func (this *Server) GetClient(id string) *Client {
	if c, ok := this.clients[id]; ok {
		return c
	}
	return nil
}

func (this *Server) IsSecure() bool { return this.secure }

func (this *Server) Close() {
	this.Info("Shutting down...")

	if this.clients != nil {
		for id, _ := range this.clients {
			this.clients[id].Close()
		}
	}

	close(this.Messages)

	if this.conn != nil {
		this.rwm.Lock()
		this.conn.Close()
		this.conn = nil
		this.rwm.Unlock()
		time.Sleep(1e9)
	}
}

func (this *Server) Open(addr string) (err os.Error) {
	if this.conn != nil {
		return
	}

	var tcpaddr *net.TCPAddr
	if tcpaddr, err = net.ResolveTCPAddr(addr); err != nil {
		return
	}

	this.rwm.Lock()
	if this.conn, err = net.ListenTCP("tcp", tcpaddr); err != nil {
		this.rwm.Unlock()
		return
	}
	this.rwm.Unlock()

	this.Info("Listening on: %v", this.conn.Addr())
	this.Info("Max clients: %d", this.maxclients)

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
				this.Messages <- lib.NewClientDisconnected(addr)
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
	var client *net.TCPConn

loop:
	for this.conn != nil {
		if client, err = this.conn.AcceptTCP(); err != nil {
			if this.conn == nil {
				return
			}

			this.Error("%T %v", err, err)
			continue loop
		}

		if len(this.clients) >= this.maxclients {
			msg := lib.NewMaxClientsReached(client.RemoteAddr())
			_ = msg.Write(client)
			client.Close()
			continue loop
		}

		go this.process(client)
	}
}

func (this *Server) process(conn *net.TCPConn) {
	var err os.Error

	// Send server version. Client can decide if it supports this server
	// version. If not, it should close the connection and go away.
	if _, err = conn.Write([]uint8{lib.MTServerVersion, ServerVersion}); err != nil {
		this.Error("%s %v", conn.RemoteAddr(), err)
		return
	}

	in := make([]byte, len(ServerName))

	// Make sure we have a relevant connection. Eg: it sends us the 'Magic
	// handshake'. if not, it's likely a connection that made a wrong turn
	// somewhere and doesn't belong here.
	if _, err = conn.Read(in); err != nil {
		this.Error("%s %v", conn.RemoteAddr(), err)
		return
	}

	if string(in) != ServerName {
		conn.Close()
		return
	}

	// Connection is relevant. We need to create a proper Client instance and
	// get rid of any clients from this source we already have. This can happen
	// when the client has disconnected/timed out for wwhatever reason and wants
	// to resume its session.
	id := conn.RemoteAddr().String()

	// Unlikely to happen, but you never know. If so, consider it a questionable
	// source and discard.
	if id == "" {
		conn.Close()
		return
	}

	// If client already exists, we have a reconnect. Close the old one.
	if _, ok := this.clients[id]; ok {
		this.clients[id].Close()
	}

	conn.SetTimeout(ClientTimeout * 6e10)

	// Store new client.
	client := newClient(conn, this.clientclosed, this.Messages)
	this.rwm.Lock()
	this.clients[id] = client
	this.rwm.Unlock()
	this.Messages <- lib.NewClientConnected(conn.RemoteAddr())

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
