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
