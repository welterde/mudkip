package lib

import "bytes"
import "net"
import "io"
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

func (this *Register) Write(w io.Writer) (err os.Error) {
	var d []byte
	buf := bytes.NewBuffer(d)
	buf.WriteByte(MTRegister)
	buf.WriteString(this.Username)
	buf.WriteByte(0x00)
	buf.WriteString(this.Password)
	buf.WriteByte(0x00)
	_, err = w.Write(buf.Bytes())
	return
}
