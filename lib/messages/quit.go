package lib

import "net"
import "io"
import "os"

type Quit struct {
	sender net.Addr
}

func NewQuit(sender net.Addr) Message {
	m := new(Quit)
	m.sender = sender
	return m
}

func (this *Quit) Id() uint8        { return MTQuit }
func (this *Quit) Sender() net.Addr { return this.sender }

func (this *Quit) Read(r io.Reader) (err os.Error) {
	// Nothing to read
	return
}

func (this *Quit) Write(w io.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTQuit})
	return
}
