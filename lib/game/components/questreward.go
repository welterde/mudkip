package lib

type QuestReward struct {
	Id    int64
	Type  ItemType
	Item  int64
	Count int
}

func NewQuestReward() *QuestReward {
	return new(QuestReward)
}
