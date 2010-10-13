package lib

import "net"
import "io"
import "bufio"
import "os"

type ServerVersion struct {
	sender  net.Addr
	Version uint8
}

func NewServerVersion(sender net.Addr) *ServerVersion {
	m := new(ServerVersion)
	m.sender = sender
	return m
}

func (this *ServerVersion) Id() uint8        { return MTServerVersion }
func (this *ServerVersion) Sender() net.Addr { return this.sender }

func (this *ServerVersion) Read(r *bufio.Reader) (err os.Error) {
	this.Version, err = r.ReadByte()
	return
}

func (this *ServerVersion) Write(w io.Writer) (err os.Error) {
	_, err = w.Write([]byte{MTServerVersion, this.Version})
	return
}
