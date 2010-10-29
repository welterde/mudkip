package lib

// A group can be anything from a clan, guild or simple fellowship of friends.
type Group struct {
	Id          int64
	Name        string
	Description string
}

func NewGroup() *Group {
	return new(Group)
}
