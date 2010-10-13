package lib

import "os"

// These functions are for convenience purposes. They allows us to (de)serialize
// object data in a uniform fashion, controlled by this library. The datastore
// implementations do not have to know about the internals of each object this
// way. If we ever change or add/remove any objects, the datastores do not have
// to be modified. Only these two functions.

// Serialize an object into a compressed bitstream, ready for storage.
func Serialize(obj Object) (data []byte, err os.Error) {
	return
}

// Deserialize an object from the given bitstream.
func Deserialize(id uint16, objtype uint8, data []byte) (obj Object, err os.Error) {
	return
}
