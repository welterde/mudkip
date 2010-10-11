package main

import "os"
import "net"
import "sync"
import "time"
import "mudkip/lib"

type Client struct {
	Messages chan lib.Message
	conn     *net.TCPConn
	rwm      *sync.RWMutex
	secure   bool
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

	var tcpaddr *net.TCPAddr
	if tcpaddr, err = net.ResolveTCPAddr(addr); err != nil {
		return
	}

	this.rwm.Lock()
	if this.conn, err = net.DialTCP("tcp", nil, tcpaddr); err != nil {
		this.rwm.Unlock()
		return
	}
	this.rwm.Unlock()

	this.conn.SetTimeout(12e10) // 2 minutes

	// Announce that we are a relevant connection. eg: we are here to use the
	// mudkip server and not just some random connection which made a wrong
	// turn somewhere.
	this.conn.Write([]byte("MUDKIP"))

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
		if msg, err = lib.ReadMessage(this.conn, this.conn.RemoteAddr()); err != nil {
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
