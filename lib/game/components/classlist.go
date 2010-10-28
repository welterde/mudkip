package lib

type ClassList []*Class

func NewClassList() ClassList {
	return make([]*Class, 0, 16)
}

func (this *ClassList) Add(v *Class) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Class, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *ClassList) Remove(v *Class) {
	this.RemoveId(v.Id)
}

func (this *ClassList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *ClassList) Clear() {
	*this = make([]*Class, 0, 16)
}

func (this ClassList) Len() int {
	return len(this)
}

func (this ClassList) Find(id int64) *Class {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this ClassList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
