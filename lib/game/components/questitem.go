package lib

type QuestItem struct {
	Id          int64
	Name        string
	Description string
}

func NewQuestItem() *QuestItem {
	return new(QuestItem)
}
