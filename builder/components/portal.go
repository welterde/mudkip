package builder

// Directions
type Direction string

const (
	DirNorth     Direction = "n"
	DirSouth     Direction = "s"
	DirEast      Direction = "e"
	DirWest      Direction = "w"
	DirNorthEast Direction = "ne"
	DirNorthWest Direction = "nw"
	DirSouthEast Direction = "se"
	DirSouthWest Direction = "sw"
)

type Portal struct {
	Dir  Direction
	Zone uint16
}

func NewPortal(dir Direction, val uint16) *Portal {
	v := new(Portal)
	v.Dir = dir
	v.Zone = val
	return v
}
