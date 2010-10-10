package lib

import "net"
import "io"
import "os"

// Builtin message type IDs. These are the bytes sent over a connection and
// are used to identify a specific message on the receiving end. The IDs are
// uint8 values. 0-54 is reserved for internal use. The rest can be assigned as
// custom Message types.
const (
	MTServerVersion uint8 = iota
	MTMaxClientsReached
	MTClientConnected
	MTClientDisconnected
	MTQuit

	// Custom message types should start here: MTMax, MTMax+1, MTMax+2 etc.
	// There is room for 200 custom message types.
	MTMax = 55
)

// A generic Message type. Any message should implement this interface.
type Message interface {
	Id() uint8
	Sender() net.Addr
	Read(io.Reader) os.Error
	Write(io.Writer) os.Error
}

// This function should construct an empty but initialized version of a
// message. It is identified by its Message ID.
type MessageBuilder func(sender net.Addr) Message

var ErrUnknownMessage os.Error

// This contains handlers for processing of various messages. You should add
// your own message types to this map if you want them to be processed.
var Messages map[uint8]MessageBuilder

func init() {
	ErrUnknownMessage = os.NewError("Unknown message type")
	Messages = make(map[uint8]MessageBuilder)

	Messages[MTServerVersion] = NewServerVersion
	Messages[MTMaxClientsReached] = NewMaxClientsReached
	Messages[MTClientConnected] = NewClientConnected
	Messages[MTClientDisconnected] = NewClientDisconnected
	Messages[MTQuit] = NewQuit
}

func ReadMessage(r io.Reader, sender net.Addr) (msg Message, err os.Error) {
	data := make([]byte, 128)

	// Read message type
	if _, err = r.Read(data[0:1]); err != nil {
		return
	}

	if builder, ok := Messages[data[0]]; ok {
		msg = builder(sender)
		err = msg.Read(r)
		return
	}

	return nil, ErrUnknownMessage
}
