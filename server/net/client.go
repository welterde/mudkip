package main

import "net"
import "sync"
import "os"
import "mudkip/lib"

type Client struct {
	conn     net.Conn
	rwm      *sync.RWMutex
	closing  chan net.Addr
	messages chan lib.Message
}

func newClient(conn net.Conn, closing chan net.Addr, messages chan lib.Message) *Client {
	c := new(Client)
	c.rwm = new(sync.RWMutex)
	c.closing = closing
	c.messages = messages
	c.conn = conn
	return c
}

func (this *Client) Send(msg lib.Message) (err os.Error) {
	err = msg.Write(this.conn)
	return
}

func (this *Client) Run() {
	var err os.Error
	var msg lib.Message
	var ok bool

	for this.conn != nil {
		if msg, err = lib.ReadMessage(this.conn, this.conn.RemoteAddr()); err != nil {
			if _, ok = err.(*net.OpError); ok || err == os.EOF {
				this.Close()
				return
			} else {
				em := lib.NewError(this.conn.RemoteAddr())
				em.FromError(err)
				em.Write(this.conn)
				continue
			}
		}

		if closed(this.messages) {
			this.Close()
			return
		}

		this.messages <- msg
	}
}

func (this *Client) Close() {
	if this.closing != nil && !closed(this.closing) {
		this.closing <- this.conn.RemoteAddr()

		this.rwm.Lock()
		this.closing = nil
		this.rwm.Unlock()
	}

	if this.conn != nil {
		this.rwm.Lock()
		this.conn.Close()
		this.conn = nil
		this.rwm.Unlock()
	}
}
