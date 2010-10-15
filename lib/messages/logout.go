package lib

import "net"
import "bufio"
import "os"

type Logout struct {
	sender  net.Addr
}

func NewLogout(sender net.Addr) *Logout {
	m := new(Logout)
	m.sender = sender
	return m
}

func (this *Logout) Id() uint8        { return MTLogout }
func (this *Logout) Sender() net.Addr { return this.sender }

func (this *Logout) Read(r *bufio.Reader) (err os.Error) {
	return
}

func (this *Logout) Write(w *bufio.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTLogout})
	return
}
