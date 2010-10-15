package lib

import "net"
import "bufio"
import "os"

type Ok struct {
	sender net.Addr
}

func NewOk(sender net.Addr) *Ok {
	m := new(Ok)
	m.sender = sender
	return m
}

func (this *Ok) Id() uint8        { return MTOk }
func (this *Ok) Sender() net.Addr { return this.sender }

func (this *Ok) Read(r *bufio.Reader) (err os.Error) {
	// Nothing to read
	return
}

func (this *Ok) Write(w *bufio.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTOk})
	return
}
