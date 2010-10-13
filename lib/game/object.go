package lib

import "bufio"
import "os"

// Object types
const (
	OTWorld uint8 = iota
	OTZone
	OTCharacter
)

// Field size limits in Unicode code points (not in bytes!)
const (
	LimitName        = 256
	LimitPassword    = 256
	LimitDescription = 4096
)

// Generic game object. Everything in this game should implement this interface.
type Object interface {
	Type() uint8 // Unique object type

	Id() uint16   // Unique datastore ID
	SetId(uint16) // set Unique datastore ID

	Name() string   // Display name of object
	SetName(string) // Set display name

	Description() string   // Display description of object
	SetDescription(string) // Set description

	Pack(*bufio.Writer) os.Error   // Pack object contents into a bit stream
	Unpack(*bufio.Reader) os.Error // Unpack object contents from bit stream
}
