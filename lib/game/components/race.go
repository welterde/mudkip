package lib

type Race struct {
	Id          int64
	Name        string
	Description string
	StatBonus   *Stats
}

func NewRace() *Race {
	v := new(Race)
	v.StatBonus = NewStats()
	return v
}
