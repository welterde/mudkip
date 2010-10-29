package lib

type QuestRewardList []*QuestReward

func NewQuestRewardList() QuestRewardList {
	return make([]*QuestReward, 0, 16)
}

func (this *QuestRewardList) Add(v *QuestReward) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*QuestReward, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *QuestRewardList) Remove(v *QuestReward) {
	this.RemoveId(v.Id)
}

func (this *QuestRewardList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *QuestRewardList) Clear() {
	*this = make([]*QuestReward, 0, 16)
}

func (this QuestRewardList) Len() int {
	return len(this)
}

func (this QuestRewardList) Get(id int64) *QuestReward {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this QuestRewardList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
