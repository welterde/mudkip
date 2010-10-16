package builder

// A character is either a player or an NPC. This counts all 'living' entities, 
// including the happy bunny rabbits hopping around town. They to can be
// formidable adversaries if you have just enjoyed a bit too much booze.
type Character struct {
	Name        string
	Description string
	Title       string
	Level       int
	Group       *Group
	Class       *Class
	Race        *Race
	Stats       Stats
}

func NewCharacter() *Character {
	v := new(Character)
	v.Stats = NewStats()
	v.Level = 1
	return v
}
