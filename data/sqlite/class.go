package store

import "os"
import "mudkip/lib"

func (this *Store) GetClass(int64) (*lib.Class, os.Error) {
	return nil, nil
}

func (this *Store) SetClass(*lib.Class) os.Error {
	return nil
}
