package lib

type CurrencyList []*Currency

func NewCurrencyList() CurrencyList {
	return make([]*Currency, 0, 16)
}

func (this *CurrencyList) Add(v *Currency) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Currency, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *CurrencyList) Remove(v *Currency) {
	this.RemoveId(v.Id)
}

func (this *CurrencyList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *CurrencyList) Clear() {
	*this = make([]*Currency, 0, 16)
}

func (this CurrencyList) Len() int {
	return len(this)
}

func (this CurrencyList) Find(id int64) *Currency {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this CurrencyList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
