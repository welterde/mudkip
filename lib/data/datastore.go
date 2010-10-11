package lib

import "os"

// This interface is the front-end for a user-selectable database backend.
type DataStore interface {
	Open(params map[string]string) os.Error
	Close()

	GetObject(objtype uint8, id uint16) (Object, os.Error)
	SetObject(Object) os.Error
}

// Represents a function which initializes a new instance of a given datastore implementation.
type StoreBuilder func() DataStore

// This map contains all registered datastores.
var storebuilder StoreBuilder

// Register a new Datastore. This is used by Datastore implementations to make
// themselves known to us. They should typically call this in their init()
// function.
func SetDataStore(builder StoreBuilder) {
	storebuilder = builder
}

// Fetch a datastore that is currently registered. Returns nil if we have none.
func GetStore() DataStore {
	if storebuilder == nil {
		return nil
	}
	return storebuilder()
}
