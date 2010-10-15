package lib

import "net"
import "bufio"
import "os"

type EnterZone struct {
	sender net.Addr
	Zone   uint16
}

func NewEnterZone(sender net.Addr) *EnterZone {
	m := new(EnterZone)
	m.sender = sender
	return m
}

func (this *EnterZone) Id() uint8        { return MTEnterZone }
func (this *EnterZone) Sender() net.Addr { return this.sender }

func (this *EnterZone) Read(r *bufio.Reader) (err os.Error) {
	d := make([]uint8, 2)

	if _, err = r.Read(d); err != nil {
		return
	}

	this.Zone = uint16(d[0]) | uint16(d[1])<<8
	return
}

func (this *EnterZone) Write(w *bufio.Writer) (err os.Error) {
	_, err = w.Write([]uint8{MTEnterZone, uint8(this.Zone), uint8(this.Zone) >> 8})
	return
}
