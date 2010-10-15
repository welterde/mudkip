package main

import "net"
import "sync"
import "os"
import "mudkip/lib"

type Client struct {
	conn      net.Conn
	rwm       *sync.RWMutex
	closing   chan net.Addr
	onMessage MessageHandler
	ack       lib.Message
}

func newClient(conn net.Conn, closing chan net.Addr, mh MessageHandler) *Client {
	c := new(Client)
	c.rwm = new(sync.RWMutex)
	c.closing = closing
	c.conn = conn
	c.onMessage = mh
	c.ack = new(lib.Ok) // just cache it. No need to reallocate all the time.
	return c
}

func (this *Client) Ack() {
	lib.WriteMessage(this.conn, this.ack)
}

func (this *Client) Err(err os.Error) {
	msg := new(lib.Error)
	msg.FromError(err)
	lib.WriteMessage(this.conn, msg)
}

func (this *Client) Send(msg lib.Message) {
	lib.WriteMessage(this.conn, msg)
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
				this.Err(err)
				continue
			}
		}

		this.onMessage(msg)
	}
}

func (this *Client) Close() {
	this.rwm.Lock()
	defer this.rwm.Unlock()

	if this.closing != nil && !closed(this.closing) {
		this.closing <- this.conn.RemoteAddr()
		this.closing = nil
	}

	this.onMessage = nil

	if this.conn != nil {
		this.conn.Close()
		this.conn = nil
	}
}
