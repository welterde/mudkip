package builder

type World struct {
	Name        string
	Description string
	Logo        string
	Zones       []*Zone
	Characters  []*Character
	Race        []*Race
	Classes     []*Class
}
