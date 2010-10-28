package lib

type ArmorList []*Armor

func NewArmorList() ArmorList {
	return make([]*Armor, 0, 16)
}

func (this *ArmorList) Add(v *Armor) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Armor, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *ArmorList) Remove(v *Armor) {
	this.RemoveId(v.Id)
}

func (this *ArmorList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *ArmorList) Clear() {
	*this = make([]*Armor, 0, 16)
}

func (this ArmorList) Len() int {
	return len(this)
}

func (this ArmorList) Find(id int64) *Armor {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this ArmorList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
