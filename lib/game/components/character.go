package lib

// Standings
const (
	Neutral uint8 = iota
	Friendly
	Enemy
)

// A character is either a player or an NPC. This counts all 'living' entities, 
// including the happy bunny rabbits hopping around town. They to can be
// formidable adversaries if you have just enjoyed a bit too much booze.
type Character struct {
	Id          int64
	Name        string
	Description string
	Title       string
	Level       int
	Group       int64
	Class       int64
	Race        int64
	Zone        int64
	BankRoll    int64
	Standing    uint8
	Stats       Stats
}

func NewCharacter() *Character {
	v := new(Character)
	v.Standing = Neutral
	v.Group = -1
	v.Class = -1
	v.Race = -1
	v.Zone = -1
	v.BankRoll = 0
	v.Level = 1
	v.Stats = NewStats()
	return v
}
