package store

import "os"
import "mudkip/lib"

func (this *Store) GetInventory(int64) (*lib.Inventory, os.Error) {
	return nil, nil
}

func (this *Store) SetInventory(*lib.Inventory) os.Error {
	return nil
}
