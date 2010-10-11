package main

import "os"
import "net"
import "io"
import "sync"
import "time"
import "crypto/tls"
import "crypto/rand"
import "mudkip/lib"

type Client struct {
	Messages chan lib.Message
	conn     io.ReadWriteCloser
	rwm      *sync.RWMutex
	secure   bool
	addr     net.Addr
}

func NewClient(secure bool) *Client {
	c := new(Client)
	c.Messages = make(chan lib.Message, 8)
	c.rwm = new(sync.RWMutex)
	c.secure = secure
	return c
}

// Close the client and it's associated connections.
func (this *Client) Close() {
	close(this.Messages)

	if this.conn != nil {
		this.rwm.Lock()
		this.conn.Close()
		this.conn = nil
		this.rwm.Unlock()
		time.Sleep(1e9)
	}
}

// Open a connection to the specified server address.
func (this *Client) Open(addr string) (err os.Error) {
	if this.conn != nil {
		return
	}

	if this.addr, err = net.ResolveTCPAddr(addr); err != nil {
		return
	}

	var tcp *net.TCPConn
	if tcp, err = net.DialTCP("tcp", nil, this.addr.(*net.TCPAddr)); err != nil {
		return
	}

	tcp.SetTimeout(12e10) // 2 minutes

	this.rwm.Lock()
	this.addr = tcp.RemoteAddr()

	if this.secure {
		// FIXME: Half-arsed SSL implementation. Does not verify certificate.
		cf := new(tls.Config)
		cf.Rand = rand.Reader
		cf.Time = time.Nanoseconds
		this.conn = tls.Client(tcp, cf)
	} else {
		this.conn = tcp
	}

	this.rwm.Unlock()

	// Announce that we are a relevant connection. eg: we are here to use the
	// mudkip server and not just some random connection which made a wrong
	// turn somewhere.
	if _, err = this.conn.Write([]byte("MUDKIP")); err != nil {
		return
	}

	go this.poll()
	return
}

// Send a message to the server.
func (this *Client) Send(msg lib.Message) (err os.Error) {
	err = msg.Write(this.conn)
	return
}

func (this *Client) poll() {
	var err os.Error
	var msg lib.Message

	for this.conn != nil {
		if msg, err = lib.ReadMessage(this.conn, this.addr); err != nil {
			this.Close()
			return
		}

		if closed(this.Messages) {
			this.Close()
			return
		}

		this.Messages <- msg
	}
}
