package lib

import "net"
import "io"
import "os"

type ServerVersion struct {
	sender  net.Addr
	Version uint8
}

func NewServerVersion(sender net.Addr) Message {
	m := new(ServerVersion)
	m.sender = sender
	return m
}

func (this *ServerVersion) Id() uint8        { return MTServerVersion }
func (this *ServerVersion) Sender() net.Addr { return this.sender }

func (this *ServerVersion) Read(r io.Reader) (err os.Error) {
	data := make([]byte, 1)

	if _, err = r.Read(data); err != nil {
		return
	}

	this.Version = data[0]
	return
}

func (this *ServerVersion) Write(w io.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTServerVersion, this.Version})
	return
}
