package lib

import "net"
import "io"
import "bufio"
import "os"

// Builtin message type IDs. These are the bytes sent over a connection and
// are used to identify a specific message on the receiving end. The IDs are
// uint8 values.
const (
	MTError uint8 = iota
	MTOk
	MTServerVersion
	MTMaxClientsReached
	MTClientConnected
	MTClientDisconnected
	MTLogin
	MTLogout
	MTRegister
	MTEnterZone
	MTLeaveZone
)

// A generic Message type. Any message should implement this interface.
type Message interface {
	Id() uint8
	Sender() net.Addr
	Read(*bufio.Reader) os.Error
	Write(*bufio.Writer) os.Error
}

func ReadMessage(r io.Reader, s net.Addr) (msg Message, err os.Error) {
	data := make([]byte, 1)
	if _, err = r.Read(data); err != nil {
		return
	}

	switch data[0] {
	case MTError:
		msg = NewError(s)
	case MTOk:
		msg = NewOk(s)
	case MTServerVersion:
		msg = NewServerVersion(s)
	case MTMaxClientsReached:
		msg = NewMaxClientsReached(s)
	case MTClientConnected:
		msg = NewClientConnected(s)
	case MTClientDisconnected:
		msg = NewClientDisconnected(s)
	case MTLogin:
		msg = NewLogin(s)
	case MTLogout:
		msg = NewLogout(s)
	case MTRegister:
		msg = NewRegister(s)
	case MTEnterZone:
		msg = NewEnterZone(s)
	case MTLeaveZone:
		msg = NewLeaveZone(s)
	default:
		return nil, ErrUnknownMessage
	}

	err = msg.Read(bufio.NewReader(r))
	return
}

func WriteMessage(w io.Writer, msg Message) (err os.Error) {
	return msg.Write(bufio.NewWriter(w))
}
