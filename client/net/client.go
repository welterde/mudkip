package main

import "os"
import "net"
import "io"
import "sync"
import "time"
import "strings"
import "crypto/tls"
import "mudkip/lib"

type Client struct {
	Messages chan lib.Message
	conn     io.ReadWriteCloser
	rwm      *sync.RWMutex
	addr     net.Addr
	config   *Config
}

func NewClient(config *Config) *Client {
	c := new(Client)
	c.Messages = make(chan lib.Message, 8)
	c.rwm = new(sync.RWMutex)
	c.config = config
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

	this.rwm.Lock()
	this.config = nil
	this.rwm.Unlock()
}

// Open a connection to the specified server address.
func (this *Client) Open() (err os.Error) {
	if this.conn != nil {
		return
	}

	if this.config.Secure {
		this.rwm.Lock()
		if this.conn, err = tls.Dial("tcp", "", this.config.Server); err != nil {
			this.rwm.Unlock()
			return
		}
		this.rwm.Unlock()

		addr := this.config.Server
		if idx := strings.LastIndex(addr, ":"); idx != -1 {
			if idx > strings.LastIndex(addr, "]") { // ipv6
				addr = addr[0:idx]
			}
		}

		if err = this.conn.(*tls.Conn).VerifyHostname(addr); err != nil {
			return err
		}

		this.rwm.Lock()
		this.addr = this.conn.(*tls.Conn).RemoteAddr()
		this.rwm.Unlock()
	} else {
		this.rwm.Lock()
		if this.conn, err = net.Dial("tcp", "", this.config.Server); err != nil {
			this.rwm.Unlock()
			return
		}

		this.addr = this.conn.(*net.TCPConn).RemoteAddr()
		this.rwm.Unlock()

		this.conn.(*net.TCPConn).SetTimeout(12e10) // 2 minutes
	}

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
