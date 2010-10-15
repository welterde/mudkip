package lib

import "net"
import "bufio"
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

func (this *MaxClientsReached) Read(r *bufio.Reader) (err os.Error) {
	// Nothing to read
	return
}

func (this *MaxClientsReached) Write(w *bufio.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTMaxClientsReached})
	return
}
