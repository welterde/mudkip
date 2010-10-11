package lib

import "net"
import "io"
import "os"

type ClientDisconnected struct {
	sender net.Addr
}

func NewClientDisconnected(sender net.Addr) *ClientDisconnected {
	m := new(ClientDisconnected)
	m.sender = sender
	return m
}

func (this *ClientDisconnected) Id() uint8        { return MTClientDisconnected }
func (this *ClientDisconnected) Sender() net.Addr { return this.sender }

func (this *ClientDisconnected) Read(r io.Reader) (err os.Error) {
	// This message will never be sent across the wire.
	return
}

func (this *ClientDisconnected) Write(w io.Writer) (err os.Error) {
	// This message will never be sent across the wire.
	return
}
