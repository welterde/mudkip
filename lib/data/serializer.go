package lib

import "os"
import "bytes"
import "bufio"
import "compress/gzip"

// These functions are for convenience purposes. They allows us to (de)serialize
// object data in a uniform fashion, controlled by this library. The datastore
// implementations do not have to know about the internals of each object this
// way. If we ever change or add/remove any objects, the datastores do not have
// to be modified. Each object implementing the lib.Object interface, has a 
// Pack() and Unpack() method which actually does the (de)serialization. These
// two functions simply forward their call to the appropriate object's Pack and
// Unpack routines.

// Serialize an object into a compressed bitstream, ready for storage.
func Serialize(obj Object) ([]byte, os.Error) {
	var d []byte
	var err os.Error
	var cmp *gzip.Compressor

	buf := bytes.NewBuffer(d)
	if cmp, err = gzip.NewWriterLevel(buf, gzip.BestCompression); err != nil {
		return nil, err
	}

	w := bufio.NewWriter(cmp)
	if err = obj.Pack(w); err != nil {
		return nil, err
	}

	w.Flush()
	cmp.Close()

	return buf.Bytes(), err
}

// Deserialize an object from the given bitstream.
func Deserialize(id uint16, objtype uint8, data []byte) (obj Object, err os.Error) {
	var cmp *gzip.Decompressor

	switch objtype {
	case OTWorld:
		obj = NewWorld()
	case OTZone:
		obj = NewZone()
	case OTCharacter:
		obj = NewCharacter()
	default:
		return nil, ErrUnknownObject
	}

	buf := bytes.NewBuffer(data)
	if cmp, err = gzip.NewReader(buf); err != nil {
		return nil, err
	}

	err = obj.Unpack(bufio.NewReader(cmp))
	return
}
