package lib

type ArmorType uint8

// Armor types. Define where a specific armor type should go: legs, chest,
// arms, etc. Do not change the order of these values, since this will change
// the meaning of any allready defined armor pieces.
const (
	Helmet ArmorType = iota
	ShoulderPads
	Cape
	Chest
	Arms
	Gloves
	Belt
	Legs
	Boots
	Wrists
	Necklace
	Ring
)

type Armor struct {
	Id          int64
	Name        string
	Description string
	Type        ArmorType
}

func NewArmor() *Armor {
	v := new(Armor)
	return v
}
