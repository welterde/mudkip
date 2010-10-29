package store

import "os"
import "mudkip/lib"

func (this *Store) GetCurrency(int64) (*lib.Currency, os.Error) {
	return nil, nil
}

func (this *Store) SetCurrency(*lib.Currency) os.Error {
	return nil
}
