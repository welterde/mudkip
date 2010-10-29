package store

import "os"
import "mudkip/lib"

func (this *Store) GetCharacter(int64) (*lib.Character, os.Error) {
	return nil, nil
}

func (this *Store) SetCharacter(*lib.Character) os.Error {
	return nil
}
