package lib

type WeaponList []*Weapon

func NewWeaponList() WeaponList {
	return make([]*Weapon, 0, 16)
}

func (this *WeaponList) Add(v *Weapon) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Weapon, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *WeaponList) Remove(v *Weapon) {
	this.RemoveId(v.Id)
}

func (this *WeaponList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *WeaponList) Clear() {
	*this = make([]*Weapon, 0, 16)
}

func (this WeaponList) Len() int {
	return len(this)
}

func (this WeaponList) Get(id int64) *Weapon {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this WeaponList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
