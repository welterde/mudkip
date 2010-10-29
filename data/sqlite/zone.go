package store

import "os"
import "mudkip/lib"

func (this *Store) GetZone(int64) (*lib.Zone, os.Error) {
	return nil, nil
}

func (this *Store) SetZone(*lib.Zone) os.Error {
	return nil
}
