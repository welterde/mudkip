package lib

import "os"

// This interface is the front-end for a user-selectable database backend.
type DataStore interface {
	Open(params map[string]string) os.Error
	Close()

	GetObject(objtype uint8, id uint16) (Object, os.Error)
	SetObject(Object) os.Error
}
