package lib

import "bytes"
import "net"
import "io"
import "os"

// Login message. Username and password can each not exceed 50 bytes.
// That includes bytes occupied by multibyte characters.
type Login struct {
	sender   net.Addr
	Username string
	Password string
}

func NewLogin(sender net.Addr) *Login {
	m := new(Login)
	m.sender = sender
	return m
}

func (this *Login) Id() uint8        { return MTLogin }
func (this *Login) Sender() net.Addr { return this.sender }

func (this *Login) Read(r io.Reader) (err os.Error) {
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

func (this *Login) Write(w io.Writer) (err os.Error) {
	var d []byte
	buf := bytes.NewBuffer(d)
	buf.WriteByte(MTLogin)
	buf.WriteByte(uint8(len(this.Username)))
	buf.WriteString(this.Username)
	buf.WriteByte(uint8(len(this.Password)))
	buf.WriteString(this.Password)
	_, err = w.Write(buf.Bytes())
	return
}
