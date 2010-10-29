package store

import "os"
import "mudkip/lib"

func (this *Store) GetGroup(int64) (*lib.Group, os.Error) {
	return nil, nil
}

func (this *Store) SetGroup(*lib.Group) os.Error {
	return nil
}
