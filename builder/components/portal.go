package builder

type Direction string

// Directions
const (
	DirNorth      Direction = "n"
	DirSouth      Direction = "s"
	DirEast       Direction = "e"
	DirWest       Direction = "w"
	DirNorthEast  Direction = "ne"
	DirNorthWest  Direction = "nw"
	DirSouthEast  Direction = "se"
	DirSouthWest  Direction = "sw"
	MaxDirections = 8
)

type Portal struct {
	Dir  Direction
	Zone int
}

func NewPortal() *Portal {
	return new(Portal)
}
