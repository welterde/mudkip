package lib

import "net"
import "io"
import "os"

type Error struct {
	sender net.Addr
	Errno  uint8
}

func NewError(sender net.Addr) *Error {
	m := new(Error)
	m.sender = sender
	return m
}

func (this *Error) Id() uint8        { return MTError }
func (this *Error) Sender() net.Addr { return this.sender }

func (this *Error) Read(r io.Reader) (err os.Error) {
	data := make([]byte, 1)

	if _, err = r.Read(data); err != nil {
		return
	}

	this.Errno = data[0]
	return
}

func (this *Error) Write(w io.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTError, this.Errno})
	return
}
