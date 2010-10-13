package lib

import "io"
import "os"

// Object types
const (
	OTWorld uint8 = iota
	OTZone
	OTCharacter
)

// Generic game object. Everything in this game should implement this interface.
type Object interface {
	Type() uint8                 // Unique object type
	Id() uint16                  // Unique datastore ID
	SetId(uint16)                // set Unique datastore ID
	Name() string                // Display name of object
	Description() string         // Display description of object
	Pack(w io.Writer) os.Error   // Pack object contents into a bit stream
	Unpack(r io.Reader) os.Error // Unpack object contents from bit stream
}
