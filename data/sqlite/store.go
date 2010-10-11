package sqlite

import "os"
import "mudkip/lib"

func init() {
	// Hook up the build method for this store.
	lib.SetDataStore(New)
}

type Store struct{}

func New() lib.DataStore { return new(Store) }

func (this *Store) Open(params map[string]string) (err os.Error) {
	return
}

func (this *Store) Close() {

}

func (this *Store) GetObject(objtype uint8, id uint16) (obj lib.Object, err os.Error) {
	return
}

func (this *Store) SetObject(lib.Object) (err os.Error) {
	return
}
