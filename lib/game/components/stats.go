package lib

type Stats struct {
	Id  int64
	HP  int8 // health points
	MP  int8 // magic points (mana, or energy or rage or whatever)
	AP  int8 // Attack Points
	DEF int8 // Defense points
	AGI int8 // Agility
	STR int8 // Strength
	WIS int8 // Wisdom
	LUC int8 // Luck
	CHR int8 // Charisma
	PER int8 // Perception
}

func NewStats() *Stats {
	return new(Stats)
}
