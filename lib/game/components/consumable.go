package lib

type Consumable struct {
	Id          int64
	Name        string
	Description string
	Liquid      bool // Eat it or drink it?
	StatBonus   *Stats
}

func NewConsumable() *Consumable {
	v := new(Consumable)
	v.Liquid = true
	v.StatBonus = NewStats()
	return v
}
