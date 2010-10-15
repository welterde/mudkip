package builder

// Directions
const (
	DirNorth = "n"
	DirSouth = "s"
	DirEast  = "e"
	DirWest  = "w"
)

type Portal struct {
	Dir  string
	Zone int
}
