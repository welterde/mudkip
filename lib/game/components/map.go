package lib

import "os"

/*
This type represents a virtual map of the entire world. It contains all zones
and their portal links, generating a 2 dimensional map of the entire world.
This can be used for visual representation of the world, or a way to ensure
that we do not define multiple zones for the same map grid section. This map
has no limit. It can grow to theoretically infinite size.
*/

type Map struct{}

type MapLink struct {
	Zone int
}

// Creates a new world map from the selected zone list
func NewMap() (m *Map) {
	return new(Map)
}

func (this *Map) Clear() {

}

func (this *Map) Fill(zones []*Zone) (err os.Error) {
	return
}
