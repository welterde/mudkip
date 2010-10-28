package lib

type World struct {
	Id            int64
	Created       int64
	Name          string
	Description   string
	Logo          string
	Motd          string
	AllowRegister bool
	LevelCap      int
	DefaultZone   int64
	Zones         []*Zone
	Characters    []*Character
	Groups        []*Group
	Races         []*Race
	Classes       []*Class
	Currency      []*Currency
	Armor         []*Armor
	Weapons       []*Weapon
}

func NewWorld() *World {
	v := new(World)
	v.Zones = make([]*Zone, 0, 32)
	v.Characters = make([]*Character, 0, 32)
	v.Races = make([]*Race, 0, 8)
	v.Classes = make([]*Class, 0, 8)
	v.Currency = make([]*Currency, 0, 8)
	v.Groups = make([]*Group, 0, 8)
	v.Armor = make([]*Armor, 0, 32)
	v.Weapons = make([]*Weapon, 0, 32)
	v.AllowRegister = true
	v.DefaultZone = -1
	return v
}

// Add a new zone
func (this *World) AddZone(v *Zone) int {
	sz := len(this.Zones)

	if sz >= cap(this.Zones) {
		cp := make([]*Zone, sz, sz+32)
		copy(cp, this.Zones)
		this.Zones = cp
	}

	this.Zones = this.Zones[0 : sz+1]
	this.Zones[sz] = v
	return sz
}

// Add a new group
func (this *World) AddGroup(v *Group) int {
	sz := len(this.Groups)

	if sz >= cap(this.Groups) {
		cp := make([]*Group, sz, sz+32)
		copy(cp, this.Groups)
		this.Groups = cp
	}

	this.Groups = this.Groups[0 : sz+1]
	this.Groups[sz] = v
	return sz
}

// Add a new character
func (this *World) AddCharacter(v *Character) int {
	sz := len(this.Characters)

	if sz >= cap(this.Characters) {
		cp := make([]*Character, sz, sz+32)
		copy(cp, this.Characters)
		this.Characters = cp
	}

	this.Characters = this.Characters[0 : sz+1]
	this.Characters[sz] = v
	return sz
}

// Add a new race
func (this *World) AddRace(v *Race) int {
	sz := len(this.Races)

	if sz >= cap(this.Races) {
		cp := make([]*Race, sz, sz+8)
		copy(cp, this.Races)
		this.Races = cp
	}

	this.Races = this.Races[0 : sz+1]
	this.Races[sz] = v
	return sz
}

// Add a new class
func (this *World) AddClass(v *Class) int {
	sz := len(this.Classes)

	if sz >= cap(this.Classes) {
		cp := make([]*Class, sz, sz+8)
		copy(cp, this.Classes)
		this.Classes = cp
	}

	this.Classes = this.Classes[0 : sz+1]
	this.Classes[sz] = v
	return sz
}

// Add a new currency
func (this *World) AddCurrency(v *Currency) int {
	sz := len(this.Currency)

	if sz >= cap(this.Currency) {
		cp := make([]*Currency, sz, sz+8)
		copy(cp, this.Currency)
		this.Currency = cp
	}

	this.Currency = this.Currency[0 : sz+1]
	this.Currency[sz] = v
	return sz
}

// Add a new weapon
func (this *World) AddWeapon(v *Weapon) int {
	sz := len(this.Weapons)

	if sz >= cap(this.Weapons) {
		cp := make([]*Weapon, sz, sz+8)
		copy(cp, this.Weapons)
		this.Weapons = cp
	}

	this.Weapons = this.Weapons[0 : sz+1]
	this.Weapons[sz] = v
	return sz
}

// Add a new armor
func (this *World) AddArmor(v *Armor) int {
	sz := len(this.Armor)

	if sz >= cap(this.Armor) {
		cp := make([]*Armor, sz, sz+8)
		copy(cp, this.Armor)
		this.Armor = cp
	}

	this.Armor = this.Armor[0 : sz+1]
	this.Armor[sz] = v
	return sz
}
