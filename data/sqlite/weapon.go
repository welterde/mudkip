package store

import "os"
import "mudkip/lib"

func (this *Store) GetWeapon(int64) (*lib.Weapon, os.Error) {
	return nil, nil
}

func (this *Store) SetWeapon(*lib.Weapon) os.Error {
	return nil
}
