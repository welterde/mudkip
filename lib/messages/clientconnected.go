package lib

import "net"
import "bufio"
import "os"

type ClientConnected struct {
	sender net.Addr
}

func NewClientConnected(sender net.Addr) *ClientConnected {
	m := new(ClientConnected)
	m.sender = sender
	return m
}

func (this *ClientConnected) Id() uint8        { return MTClientConnected }
func (this *ClientConnected) Sender() net.Addr { return this.sender }

func (this *ClientConnected) Read(r *bufio.Reader) (err os.Error) {
	// This message will never be sent across the wire.
	return
}

func (this *ClientConnected) Write(w *bufio.Writer) (err os.Error) {
	// This message will never be sent across the wire.
	return
}
