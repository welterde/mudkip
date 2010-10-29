package lib

type RaceList []*Race

func NewRaceList() RaceList {
	return make([]*Race, 0, 16)
}

func (this *RaceList) Add(v *Race) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Race, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *RaceList) Remove(v *Race) {
	this.RemoveId(v.Id)
}

func (this *RaceList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *RaceList) Clear() {
	*this = make([]*Race, 0, 16)
}

func (this RaceList) Len() int {
	return len(this)
}

func (this RaceList) Get(id int64) *Race {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this RaceList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
