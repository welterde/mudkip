package lib

type QuestList []*Quest

func NewQuestList() QuestList {
	return make([]*Quest, 0, 16)
}

func (this *QuestList) Add(v *Quest) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Quest, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *QuestList) Remove(v *Quest) {
	this.RemoveId(v.Id)
}

func (this *QuestList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *QuestList) Clear() {
	*this = make([]*Quest, 0, 16)
}

func (this QuestList) Len() int {
	return len(this)
}

func (this QuestList) Get(id int64) *Quest {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this QuestList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
