package store

import "os"
import "mudkip/lib"

func (this *Store) GetQuestReward(int64) (*lib.QuestReward, os.Error) {
	return nil, nil
}

func (this *Store) SetQuestReward(*lib.QuestReward) os.Error {
	return nil
}
