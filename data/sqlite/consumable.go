package store

import "os"
import "mudkip/lib"

func (this *Store) GetConsumable(int64) (*lib.Consumable, os.Error) {
	return nil, nil
}

func (this *Store) SetConsumable(*lib.Consumable) os.Error {
	return nil
}
