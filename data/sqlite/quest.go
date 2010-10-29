package store

import "os"
import "mudkip/lib"

func (this *Store) GetQuest(int64) (*lib.Quest, os.Error) {
	return nil, nil
}

func (this *Store) SetQuest(*lib.Quest) os.Error {
	return nil
}
