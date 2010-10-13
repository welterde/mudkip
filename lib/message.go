package lib

import "net"
import "io"
import "bufio"
import "os"

// Builtin message type IDs. These are the bytes sent over a connection and
// are used to identify a specific message on the receiving end. The IDs are
// uint8 values. 0-54 is reserved for internal use. The rest can be assigned as
// custom Message types.
const (
	MTError uint8 = iota
	MTServerVersion
	MTMaxClientsReached
	MTClientConnected
	MTClientDisconnected
	MTLogin
	MTRegister

	// Custom message types should start here: MTMax, MTMax+1, MTMax+2 etc.
	// There is room for 200 custom message types.
	MTMax = 55
)

// A generic Message type. Any message should implement this interface.
type Message interface {
	Id() uint8
	Sender() net.Addr
	Read(*bufio.Reader) os.Error
	Write(io.Writer) os.Error
}

// This function should construct an empty but initialized version of a
// message. It is identified by its Message ID.
type MessageBuilder func(sender net.Addr) Message


// This contains handlers for processing of various messages. You should add
// your own message types to this map if you want them to be processed.
var Messages map[uint8]MessageBuilder

func init() {
	Messages = make(map[uint8]MessageBuilder)

	Messages[MTError] = func(s net.Addr) Message { return NewError(s) }
	Messages[MTServerVersion] = func(s net.Addr) Message { return NewServerVersion(s) }
	Messages[MTMaxClientsReached] = func(s net.Addr) Message { return NewMaxClientsReached(s) }
	Messages[MTClientConnected] = func(s net.Addr) Message { return NewClientConnected(s) }
	Messages[MTClientDisconnected] = func(s net.Addr) Message { return NewClientDisconnected(s) }
	Messages[MTLogin] = func(s net.Addr) Message { return NewLogin(s) }
	Messages[MTRegister] = func(s net.Addr) Message { return NewRegister(s) }
}

func ReadMessage(r io.Reader, sender net.Addr) (msg Message, err os.Error) {
	data := make([]byte, 1)

	// Read message type
	if _, err = r.Read(data); err != nil {
		return
	}

	if builder, ok := Messages[data[0]]; ok {
		msg = builder(sender)
		err = msg.Read(bufio.NewReader(r))
		return
	}

	return nil, ErrUnknownMessage
}
