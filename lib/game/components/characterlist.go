package lib

type CharacterList []*Character

func NewCharacterList() CharacterList {
	return make([]*Character, 0, 16)
}

func (this *CharacterList) Add(v *Character) int64 {
	sz := len(*this)

	if sz >= cap(*this) {
		cp := make([]*Character, sz, sz+16)
		copy(cp, *this)
		*this = cp
	}

	v.Id = int64(sz) + 1
	*this = (*this)[0 : sz+1]
	(*this)[sz] = v
	return v.Id
}

func (this *CharacterList) Remove(v *Character) {
	this.RemoveId(v.Id)
}

func (this *CharacterList) RemoveId(id int64) {
	idx := this.IndexOf(id)
	if idx == -1 {
		return
	}

	copy((*this)[idx:], (*this)[idx+1:])
	*this = (*this)[:len(*this)-1]
}

func (this *CharacterList) Clear() {
	*this = make([]*Character, 0, 16)
}

func (this CharacterList) Len() int {
	return len(this)
}

func (this CharacterList) Find(id int64) *Character {
	idx := this.IndexOf(id)
	if idx == -1 {
		return nil
	}
	return this[idx]
}

func (this CharacterList) IndexOf(id int64) int {
	for i, v := range this {
		if v.Id == id {
			return i
		}
	}
	return -1
}
