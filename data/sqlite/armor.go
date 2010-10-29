package store

import "os"
import "mudkip/lib"

func (this *Store) GetArmor(int64) (*lib.Armor, os.Error) {
	return nil, nil
}

func (this *Store) SetArmor(*lib.Armor) os.Error {
	return nil
}
