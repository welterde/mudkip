package builder

type Zone struct {
	Id          int
	Default     bool
	Name        string
	Description string
	Lighting    string
	Smell       string
	Sound       string
	Exits       []Portal
}

func NewZone() *Zone {
	v := new(Zone)
	v.Exits = make([]Portal, 0, 8)
	return v
}
