package lib

import "net"
import "bufio"
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

// Returns the matching os.Error from the Error.Errno.
func (this *Error) ToError() os.Error { return intToErr(this.Errno) }

// Sets Error.Errno to the appropriate value depending on the supplied os.Error
func (this *Error) FromError(err os.Error) { this.Errno = errToInt(err) }

func (this *Error) Read(r *bufio.Reader) (err os.Error) {
	this.Errno, err = r.ReadByte()
	return
}

func (this *Error) Write(w *bufio.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTError, this.Errno})
	return
}
