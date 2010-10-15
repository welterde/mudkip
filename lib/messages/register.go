package lib

import "net"
import "bufio"
import "os"

// Register message. Username and password can each not exceed 50 bytes.
// That includes bytes occupied by multibyte characters.
type Register struct {
	sender   net.Addr
	Username string
	Password string
}

func NewRegister(sender net.Addr) *Register {
	m := new(Register)
	m.sender = sender
	return m
}

func (this *Register) Id() uint8        { return MTRegister }
func (this *Register) Sender() net.Addr { return this.sender }

func (this *Register) Read(r *bufio.Reader) (err os.Error) {
	var data []byte

	if data, err = r.ReadBytes(0x00); err != nil {
		return
	}

	if len(data) > 0 {
		data = data[0 : len(data)-1]
		this.Username = string(data)
	}

	if data, err = r.ReadBytes(0x00); err != nil {
		return
	}

	if len(data) > 0 {
		data = data[0 : len(data)-1]
		this.Password = string(data)
	}

	return
}

func (this *Register) Write(w *bufio.Writer) (err os.Error) {
	w.WriteByte(MTRegister)
	w.WriteString(this.Username)
	w.WriteByte(0x00)
	w.WriteString(this.Password)
	return w.WriteByte(0x00)
}
