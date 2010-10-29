package lib

type QuestItemList []*QuestItem

func NewQuestItemList() QuestItemList {
	return make([]*QuestItem, 0, 16)
}

func (this *QuestItemList) Add(v *QuestItem) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*QuestItem, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *QuestItemList) Remove(v *QuestItem) {
	this.RemoveId(v.Id)
}

func (this *QuestItemList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *QuestItemList) Clear() {
	*this = make([]*QuestItem, 0, 16)
}

func (this QuestItemList) Len() int {
	return len(this)
}

func (this QuestItemList) Get(id int64) *QuestItem {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this QuestItemList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
