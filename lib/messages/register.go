package lib

import "bytes"
import "net"
import "io"
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

func (this *Register) Read(r io.Reader) (err os.Error) {
	data := make([]byte, 50)

	// Read length of username
	if _, err = r.Read(data[0:1]); err != nil {
		return
	}

	if data[0] > 50 {
		return ErrInvalidUsername
	}

	if _, err = r.Read(data[0:data[0]]); err != nil {
		return
	}

	this.Username = string(data[0:data[0]])

	// Read password length
	if _, err = r.Read(data[0:1]); err != nil {
		return
	}

	if data[0] > 50 {
		return ErrInvalidPassword
	}

	if _, err = r.Read(data[0:data[0]]); err != nil {
		return
	}

	this.Password = string(data[0:data[0]])
	return
}

func (this *Register) Write(w io.Writer) (err os.Error) {
	var d []byte
	buf := bytes.NewBuffer(d)
	buf.WriteByte(MTRegister)
	buf.WriteByte(uint8(len(this.Username)))
	buf.WriteString(this.Username)
	buf.WriteByte(uint8(len(this.Password)))
	buf.WriteString(this.Password)
	_, err = w.Write(buf.Bytes())
	return
}
