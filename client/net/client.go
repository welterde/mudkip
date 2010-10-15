package main

import "os"
import "net"
import "io"
import "sync"
import "strings"
import "crypto/tls"
import "mudkip/lib"

type Client struct {
	messages  chan lib.Message
	conn      io.ReadWriteCloser
	rwm       *sync.RWMutex
	addr      net.Addr
	config    *Config
}

func NewClient(config *Config) *Client {
	c := new(Client)
	c.messages = make(chan lib.Message)
	c.rwm = new(sync.RWMutex)
	c.config = config
	return c
}

func (this *Client) Messages() <-chan lib.Message { return this.messages }

// Close the client and it's associated connections.
func (this *Client) Close() {
	this.rwm.Lock()
	defer this.rwm.Unlock()

	close(this.messages)

	if this.conn != nil {
		this.conn.Close()
		this.conn = nil
	}

	this.config = nil
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

		this.addr = this.conn.(*net.TCPConn).RemoteAddr()
		this.rwm.Unlock()

		if !this.config.AcceptInvalidCert {
			addr := this.config.Server
			if idx := strings.LastIndex(addr, ":"); idx != -1 {
				if idx > strings.LastIndex(addr, "]") { // ipv6
					addr = addr[0:idx]
				}
			}

			if err = this.conn.(*tls.Conn).VerifyHostname(addr); err != nil {
				return err
			}
		}
	} else {
		this.rwm.Lock()
		if this.conn, err = net.Dial("tcp", "", this.config.Server); err != nil {
			this.rwm.Unlock()
			return
		}

		this.addr = this.conn.(*net.TCPConn).RemoteAddr()
		this.rwm.Unlock()
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
	return lib.WriteMessage(this.conn, msg)
}

func (this *Client) poll() {
	var err os.Error
	var msg lib.Message

	for this.conn != nil {
		if msg, err = lib.ReadMessage(this.conn, this.addr); err != nil {
			this.Close()
			return
		}

		this.messages <- msg
	}
}
