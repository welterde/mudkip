package lib

import "bytes"
import "net"
import "io"
import "bufio"
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

func (this *Login) Read(r *bufio.Reader) (err os.Error) {
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

func (this *Login) Write(w io.Writer) (err os.Error) {
	var d []byte
	buf := bytes.NewBuffer(d)
	buf.WriteByte(MTLogin)
	buf.WriteString(this.Username)
	buf.WriteByte(0x00)
	buf.WriteString(this.Password)
	buf.WriteByte(0x00)
	_, err = w.Write(buf.Bytes())
	return
}
