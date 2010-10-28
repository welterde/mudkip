package lib

type WeaponType uint8

// Weapon type flags - This determines what type of weapon we have. 1 handed, 2 handed,
// melee or ranged, etc
const (
	Ranged    WeaponType = 1 << iota // gun, bow
	Melee                            // sword, axe, mace, dagger, staff
	TwoHanded                        // sword, axe, mace, dagger, staff
	OneHanded                        // sword, axe, mace, dagger
	MainHand                         // sword, axe, mace, dagger
	Offhand                          // sword, axe, mace, dagger, tome or other offhand trinket like item
)

type Weapon struct {
	Id          int64
	Type        WeaponType
	Name        string
	Description string
	Damage      [2]int
}

func NewWeapon() *Weapon {
	v := new(Weapon)
	v.Type = Melee | OneHanded
	v.Damage = [2]int{0, 0}
	return v
}
