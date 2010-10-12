package main

import "os"
import "net"
import "sync"
import "time"
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
	messages      chan lib.Message
	conn          net.Listener
	lock          *sync.RWMutex
	clients       map[string]*Client
	clientclosed  chan net.Addr
	maxclients    int
	clientTimeout int64
}

func NewServer(maxclients, timeout int) *Server {
	s := new(Server)
	s.lock = new(sync.RWMutex)
	s.clients = make(map[string]*Client)
	s.messages = make(chan lib.Message, 32)
	s.clientclosed = make(chan net.Addr)
	s.maxclients = maxclients
	s.clientTimeout = int64(timeout) * 6e10
	return s
}

// Get the client associated with the given ID.
func (this *Server) GetClient(id string) *Client {
	if c, ok := this.clients[id]; ok {
		return c
	}
	return nil
}

// Channel yielding incoming messages
func (this *Server) Messages() <-chan lib.Message { return this.messages }

// Close the server. Shuts down client connections and the listening socket.
func (this *Server) Close() {
	if this.clients != nil {
		for id, _ := range this.clients {
			this.clients[id].Close()
		}
	}

	close(this.messages)

	this.lock.Lock()
	if this.conn != nil {
		this.conn.Close()
		this.conn = nil
		this.lock.Unlock()
		time.Sleep(1e9)
		return
	}

	this.lock.Unlock()
}

// Open the server and start listening
func (this *Server) Open(listenaddr string, secure bool, certfile, keyfile string) (err os.Error) {
	if this.conn != nil {
		return
	}

	this.lock.Lock()
	if this.conn, err = net.Listen("tcp", listenaddr); err != nil {
		this.lock.Unlock()
		return
	}
	this.lock.Unlock()

	if secure {
		conf := new(tls.Config)
		conf.Rand = rand.Reader
		conf.Time = time.Nanoseconds
		conf.Certificates = make([]tls.Certificate, 1)

		if conf.Certificates[0], err = tls.LoadX509KeyPair(certfile, keyfile); err != nil {
			return
		}

		this.lock.Lock()
		this.conn = tls.NewListener(this.conn, conf)
		this.lock.Unlock()
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
				this.lock.Lock()

				id = addr.String()
				if _, ok := this.clients[id]; ok {
					this.clients[id] = nil, false
				}

				this.lock.Unlock()
			}
		}

		if closed(this.clientclosed) {
			return
		}
	}
}

// Poll for incoming connections
func (this *Server) poll() {
	var err os.Error
	var client net.Conn

loop:
	for this.conn != nil {
		if client, err = this.conn.Accept(); err != nil {
			if this.conn == nil {
				return
			}
			continue loop
		}

		if len(this.clients) >= this.maxclients {
			client.Write([]uint8{lib.MTMaxClientsReached})
			client.Close()
			continue loop
		}

		go this.process(client)
	}
}

// Process client connection
func (this *Server) process(conn net.Conn) {
	var err os.Error

	// Send server version. Client can decide if it supports this server
	// version. If not, it should close the connection and go away.
	if _, err = conn.Write([]uint8{lib.MTServerVersion, ServerVersion}); err != nil {
		return
	}

	in := make([]byte, len(ServerName))

	// Make sure we have a relevant connection. Eg: it sends us the 'Magic
	// handshake'. if not, it's likely a connection that made a wrong turn
	// somewhere and doesn't belong here. Portscanners have a tendency to
	// get into this situation. They bombard us with a range of standard
	// service requests in the hopes of getting a useful response.
	if _, err = conn.Read(in); err != nil {
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
	conn.SetTimeout(this.clientTimeout)

	// If client already exists, we have a reconnect. Close the old one.
	if _, ok := this.clients[id]; ok {
		this.clients[id].Close()
	}

	// Store new client.
	client := newClient(conn, this.clientclosed, this.messages)

	this.lock.Lock()
	this.clients[id] = client
	this.lock.Unlock()
	this.messages <- lib.NewClientConnected(conn.RemoteAddr())

	// Let's get this show on the road!
	go client.Run()
}
