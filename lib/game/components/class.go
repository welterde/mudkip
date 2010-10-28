package lib

type Class struct {
	Id          int64
	Name        string
	Description string
	StatBonus   Stats
}

func NewClass() *Class {
	v := new(Class)
	v.StatBonus = NewStats()
	return v
}
