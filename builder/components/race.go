package builder

type Race struct{}

func NewRace() *Race {
	v := new(Race)
	return v
}
