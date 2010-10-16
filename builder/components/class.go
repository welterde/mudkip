package builder

type Class struct {
	Name        string
	Description string
	StatBonus   Stats
}

func NewClass() *Class {
	v := new(Class)
	v.StatBonus = NewStats()
	return v
}
