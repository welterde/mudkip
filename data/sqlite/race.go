package store

import "os"
import "mudkip/lib"

func (this *Store) GetRace(int64) (*lib.Race, os.Error) {
	return nil, nil
}

func (this *Store) SetRace(*lib.Race) os.Error {
	return nil
}
