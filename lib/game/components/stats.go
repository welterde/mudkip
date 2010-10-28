package lib

type Stats struct {
	Id  int64
	HP  uint8 // health points
	MP  uint8 // magic points (mana, or energy or rage or whatever)
	AP  uint8 // Attack Points
	DEF uint8 // Defense points
	AGI uint8 // Agility
	STR uint8 // Strength
	WIS uint8 // Wisdom
	LUC uint8 // Luck
	CHR uint8 // Charisma
	PER uint8 // Perception
}

func NewStats() *Stats {
	return new(Stats)
}
