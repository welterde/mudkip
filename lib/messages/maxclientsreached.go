package lib

import "net"
import "io"
import "os"

type MaxClientsReached struct {
	sender net.Addr
}

func NewMaxClientsReached(sender net.Addr) *MaxClientsReached {
	m := new(MaxClientsReached)
	m.sender = sender
	return m
}

func (this *MaxClientsReached) Id() uint8        { return MTMaxClientsReached }
func (this *MaxClientsReached) Sender() net.Addr { return this.sender }

func (this *MaxClientsReached) Read(r io.Reader) (err os.Error) {
	// Nothing to read
	return
}

func (this *MaxClientsReached) Write(w io.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTMaxClientsReached})
	return
}
