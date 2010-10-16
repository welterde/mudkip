package builder

type Stat uint8

const (
	HP  Stat = iota // health points
	MP              // magic points (mana, or energy or rage or whatever)
	AP              // Attack Points
	DEF             // Defense points
	AGI             // Agility
	STR             // Strength
	WIS             // Wisdom
	LUC             // Luck
	CHR             // Charisma
	PER             // Perception
)

type Stats []uint8

func NewStats() Stats {
	s := make(Stats, 10)
	s[HP] = 100
	s[MP] = 100
	return s
}
