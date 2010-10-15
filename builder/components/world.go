package builder

import "os"

type World struct {
	Name        string
	Description string
	Logo        string
	Zones       []*Zone
	Characters  []*Character
	Groups      []*Group
	Races       []*Race
	Classes     []*Class
	Currency    []*Currency
}

func NewWorld() *World {
	w := new(World)
	w.Zones = make([]*Zone, 0, 32)
	w.Characters = make([]*Character, 0, 32)
	w.Races = make([]*Race, 0, 8)
	w.Classes = make([]*Class, 0, 8)
	w.Currency = make([]*Currency, 0, 8)
	w.Groups = make([]*Group, 0, 8)
	return w
}

// This function goes through the entire data structure and finds irregularities
// in any of the components. Duplicate objects, unlinked zones, inconsistent
// bits and bobs, etc and reports them as a list of errors.
func (this *World) Sanitize() (errlist []os.Error) {
	errlist = make([]os.Error, 0, 10)

	if len(this.Name) == 0 {
		addError(&errlist, ErrNoWorldName)
	}

	if len(this.Description) == 0 {
		addError(&errlist, ErrNoWorldDescription)
	}

	if len(this.Zones) == 0 {
		addError(&errlist, ErrNoZones)
	}

	if len(this.Characters) == 0 {
		addError(&errlist, ErrNoCharacters)
	}

	if len(this.Races) == 0 {
		addError(&errlist, ErrNoRaces)
	}

	if len(this.Classes) == 0 {
		addError(&errlist, ErrNoClasses)
	}

	if len(this.Currency) == 0 {
		addError(&errlist, ErrNoCurrency)
	}

	return
}

// Add a new zone
func (this *World) AddZone(v *Zone) {
	sz := len(this.Zones)

	if sz >= cap(this.Zones) {
		cp := make([]*Zone, sz, sz+32)
		copy(cp, this.Zones)
		this.Zones = cp
	}

	this.Zones = this.Zones[0:sz+1]
	this.Zones[sz] = v
}

// Add a new group
func (this *World) AddGroup(v *Group) {
	sz := len(this.Groups)

	if sz >= cap(this.Groups) {
		cp := make([]*Group, sz, sz+32)
		copy(cp, this.Groups)
		this.Groups = cp
	}

	this.Groups = this.Groups[0:sz+1]
	this.Groups[sz] = v
}

// Add a new character
func (this *World) AddCharacter(v *Character) {
	sz := len(this.Characters)

	if sz >= cap(this.Characters) {
		cp := make([]*Character, sz, sz+32)
		copy(cp, this.Characters)
		this.Characters = cp
	}

	this.Characters = this.Characters[0:sz+1]
	this.Characters[sz] = v
}

// Add a new race
func (this *World) AddRace(v *Race) {
	sz := len(this.Races)

	if sz >= cap(this.Races) {
		cp := make([]*Race, sz, sz+8)
		copy(cp, this.Races)
		this.Races = cp
	}

	this.Races = this.Races[0:sz+1]
	this.Races[sz] = v
}

// Add a new class
func (this *World) AddClass(v *Class) {
	sz := len(this.Classes)

	if sz >= cap(this.Classes) {
		cp := make([]*Class, sz, sz+8)
		copy(cp, this.Classes)
		this.Classes = cp
	}

	this.Classes = this.Classes[0:sz+1]
	this.Classes[sz] = v
}
