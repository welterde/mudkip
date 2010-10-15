package lib

import "net"
import "bufio"
import "os"

type LeaveZone struct {
	sender net.Addr
	Zone   uint16
}

func NewLeaveZone(sender net.Addr) *LeaveZone {
	m := new(LeaveZone)
	m.sender = sender
	return m
}

func (this *LeaveZone) Id() uint8        { return MTLeaveZone }
func (this *LeaveZone) Sender() net.Addr { return this.sender }

func (this *LeaveZone) Read(r *bufio.Reader) (err os.Error) {
	d := make([]uint8, 2)

	if _, err = r.Read(d); err != nil {
		return
	}

	this.Zone = uint16(d[0]) | uint16(d[1])<<8
	return
}

func (this *LeaveZone) Write(w *bufio.Writer) (err os.Error) {
	_, err = w.Write([]uint8{MTLeaveZone, uint8(this.Zone), uint8(this.Zone) >> 8})
	return
}
