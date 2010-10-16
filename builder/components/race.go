package builder

type Race struct{
	Name        string
	Description string
}

func NewRace() *Race {
	v := new(Race)
	return v
}
