package lib

import "os"
import "sync"
import "strings"

// This interface is the front-end for a user-selectable database backend.
type DataStore interface {
	Open(params map[string]string) os.Error
	Close()

	GetObject(objtype uint8, id uint16) (Object, os.Error)
	SetObject(Object) os.Error
}

// Represents a function which initializes a new instance of a given datastore
// imlpementation.
type StoreBuilder func() DataStore

// This map contains all registered datastores.
var stores = make(map[string]StoreBuilder)
var storelock = new(sync.RWMutex)

// Register a new Datastore. This is used by Datastore implementations to make
// themselves known to us. They should typically call this in their init()
// function.
func RegisterStore(name string, builder StoreBuilder) {
	storelock.Lock()
	stores[strings.ToLower(name)] = builder
	storelock.Unlock()
}

// Fetch a datastore by the given name. Returns nil if none of the given name
// exists.
func GetStore(name string) DataStore {
	if build, ok := stores[strings.ToLower(name)]; ok {
		return build()
	}
	return nil
}

// List all available datastore names
func ListStores() []string {
	var i int

	storelock.Lock()
	list := make([]string, len(stores))
	for k, _ := range stores {
		list[i] = k
		i++
	}
	storelock.Unlock()

	return list
}
