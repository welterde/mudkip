package store

import "os"
import "mudkip/lib"

func (this *Store) GetStats(int64) (*lib.Stats, os.Error) {
	return nil, nil
}

func (this *Store) SetStats(*lib.Stats) os.Error {
	return nil
}
