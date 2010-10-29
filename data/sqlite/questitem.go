package store

import "os"
import "mudkip/lib"

func (this *Store) GetQuestItem(int64) (*lib.QuestItem, os.Error) {
	return nil, nil
}

func (this *Store) SetQuestItem(*lib.QuestItem) os.Error {
	return nil
}
