package lib

type Quest struct {
	Id          int64
	Character   int64
	Name        string
	Description string
	Rewards     QuestRewardList
}

func NewQuest() *Quest {
	v := new(Quest)
	v.Rewards = NewQuestRewardList()
	return v
}
