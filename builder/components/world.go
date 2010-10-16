package builder

type World struct {
	Name          string
	Description   string
	Logo          string
	LevelCap      int
	AllowRegister bool
	Zones         []*Zone
	Characters    []*Character
	Groups        []*Group
	Races         []*Race
	Classes       []*Class
	Currency      []*Currency
}

func NewWorld() *World {
	v := new(World)
	v.Zones = make([]*Zone, 0, 32)
	v.Characters = make([]*Character, 0, 32)
	v.Races = make([]*Race, 0, 8)
	v.Classes = make([]*Class, 0, 8)
	v.Currency = make([]*Currency, 0, 8)
	v.Groups = make([]*Group, 0, 8)
	v.AllowRegister = true
	return v
}

// This function goes through the entire data structure and finds irregularities
// in any of the components. Duplicate objects, unlinked zones, inconsistent
// bits and bobs etc and reports them as a list of builder.Error. These are not
// necessarily fatal errors. This will depend on the nature of the game being
// implemented and the wishes of the game master. These errors should just be
// considered a guide to the correct formation of a game world.
func (this *World) Sanitize() (errlist []*Error) {
	errlist = make([]*Error, 0, 10)

	if len(this.Name) == 0 {
		addError(&errlist, ErrNoObjectName, 0, this)
	}

	if len(this.Description) == 0 {
		addError(&errlist, ErrNoObjectDescription, 0, this)
	}

	if len(this.Zones) == 0 {
		addError(&errlist, ErrNoZones, 0, this)
	} else {
		var havedefault bool
		for i, v := range this.Zones {
			if len(v.Name) == 0 {
				addError(&errlist, ErrNoObjectName, i, v)
			}

			if len(v.Description) == 0 {
				addError(&errlist, ErrNoObjectDescription, i, v)
			}

			if !v.Default && len(v.Exits) == 0 {
				addError(&errlist, ErrZoneIsolated, i, v)
			}

			if v.Default {
				havedefault = true
			}
		}

		if !havedefault {
			addError(&errlist, ErrNoDefaultZone, 0, this)
		}
	}

	if len(this.Characters) == 0 {
		addError(&errlist, ErrNoCharacters, 0, this)
	} else {
		for i, v := range this.Characters {
			if len(v.Name) == 0 {
				addError(&errlist, ErrNoObjectName, i, v)
			}

			if len(v.Description) == 0 {
				addError(&errlist, ErrNoObjectDescription, i, v)
			}

			if v.Class == -1 {
				addError(&errlist, ErrNoCharacterClass, i, v)
			}

			if v.Race == -1 {
				addError(&errlist, ErrNoCharacterRace, i, v)
			}
		}
	}

	if len(this.Races) == 0 {
		addError(&errlist, ErrNoRaces, 0, this)
	} else {
		for i, v := range this.Races {
			if len(v.Name) == 0 {
				addError(&errlist, ErrNoObjectName, i, v)
			}

			if len(v.Description) == 0 {
				addError(&errlist, ErrNoObjectDescription, i, v)
			}
		}
	}

	if len(this.Classes) == 0 {
		addError(&errlist, ErrNoClasses, 0, this)
	} else {
		for i, v := range this.Classes {
			if len(v.Name) == 0 {
				addError(&errlist, ErrNoObjectName, i, v)
			}

			if len(v.Description) == 0 {
				addError(&errlist, ErrNoObjectDescription, i, v)
			}
		}
	}

	if len(this.Currency) == 0 {
		addError(&errlist, ErrNoCurrency, 0, this)
	} else {
		for i, v := range this.Currency {
			if len(v.Name) == 0 {
				addError(&errlist, ErrNoObjectName, i, v)
			}

			for j, v1 := range this.Currency {
				if i != j && v.Value == v1.Value {
					addError(&errlist, ErrDuplicateCurrencyValue, i, v)
				}
			}
		}
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

	this.Zones = this.Zones[0 : sz+1]
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

	this.Groups = this.Groups[0 : sz+1]
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

	this.Characters = this.Characters[0 : sz+1]
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

	this.Races = this.Races[0 : sz+1]
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

	this.Classes = this.Classes[0 : sz+1]
	this.Classes[sz] = v
}
