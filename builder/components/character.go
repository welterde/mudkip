package builder

type Standing uint8

// Standings
const (
	Neutral Standing = iota
	Friendly
	Enemy
)

// A character is either a player or an NPC. This counts all 'living' entities, 
// including the happy bunny rabbits hopping around town. They to can be
// formidable adversaries if you have just enjoyed a bit too much booze.
type Character struct {
	Name        string
	Description string
	Title       string
	Level       int
	Group       int
	Class       int
	Race        int
	BankRoll    int64
	Stats       Stats
	Standing    Standing
}

func NewCharacter() *Character {
	v := new(Character)
	v.Stats = NewStats()
	v.Group = -1
	v.Class = -1
	v.Race = -1
	v.Level = 1
	v.BankRoll = 0
	v.Standing = Neutral
	return v
}
