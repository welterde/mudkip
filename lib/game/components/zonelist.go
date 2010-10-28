package lib

type ZoneList []*Zone

func NewZoneList() ZoneList {
	return make([]*Zone, 0, 16)
}

func (this *ZoneList) Add(v *Zone) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Zone, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *ZoneList) Remove(v *Zone) {
	this.RemoveId(v.Id)
}

func (this *ZoneList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *ZoneList) Clear() {
	*this = make([]*Zone, 0, 16)
}

func (this ZoneList) Len() int {
	return len(this)
}

func (this ZoneList) Find(id int64) *Zone {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this ZoneList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
