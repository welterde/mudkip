package builder

type Race struct {
	Name        string
	Description string
	StatBonus   Stats
}

func NewRace() *Race {
	v := new(Race)
	v.StatBonus = NewStats()
	return v
}
