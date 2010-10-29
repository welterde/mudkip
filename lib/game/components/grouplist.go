package lib

type GroupList []*Group

func NewGroupList() GroupList {
	return make([]*Group, 0, 16)
}

func (this *GroupList) Add(v *Group) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Group, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *GroupList) Remove(v *Group) {
	this.RemoveId(v.Id)
}

func (this *GroupList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *GroupList) Clear() {
	*this = make([]*Group, 0, 16)
}

func (this GroupList) Len() int {
	return len(this)
}

func (this GroupList) Get(id int64) *Group {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this GroupList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
