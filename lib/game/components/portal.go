package lib

type Direction uint16

// Directions
const (
	North Direction = 1 << iota
	South
	East
	West
	NorthEast = North | East
	NorthWest = North | West
	SouthEast = South | East
	SouthWest = South | West
)

type Portal struct {
	Id   int64
	Zone int64
	Dir  Direction
}

func NewPortal() *Portal {
	return new(Portal)
}
