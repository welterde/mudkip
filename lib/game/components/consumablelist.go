package lib

type ConsumableList []*Consumable

func NewConsumableList() ConsumableList {
	return make([]*Consumable, 0, 16)
}

func (this *ConsumableList) Add(v *Consumable) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Consumable, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *ConsumableList) Remove(v *Consumable) {
	this.RemoveId(v.Id)
}

func (this *ConsumableList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *ConsumableList) Clear() {
	*this = make([]*Consumable, 0, 16)
}

func (this ConsumableList) Len() int {
	return len(this)
}

func (this ConsumableList) Get(id int64) *Consumable {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this ConsumableList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
