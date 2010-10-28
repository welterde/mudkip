package builder

type World struct {
	Name          string
	Description   string
	Logo          string
	AllowRegister bool
	LevelCap      int
	DefaultZone   int
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
	v.DefaultZone = -1
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
		addError(&errlist, NewError(ErrNoObjectName, 0, this))
	}

	if len(this.Description) == 0 {
		addError(&errlist, NewError(ErrNoObjectDescription, 0, this))
	}

	if this.DefaultZone == -1 {
		addError(&errlist, NewError(ErrNoDefaultZone, 0, this))
	}

	this.sanitizeZones(&errlist)
	this.sanitizeCharacters(&errlist)
	this.sanitizeRaces(&errlist)
	this.sanitizeClasses(&errlist)
	this.sanitizeCurrency(&errlist)
	return
}

func (this *World) sanitizeZones(errlist *[]*Error) {
	if len(this.Zones) == 0 {
		addError(errlist, NewError(ErrNoZones, 0, this))
		return
	}

	for i, v := range this.Zones {
		if len(v.Name) == 0 {
			addError(errlist, NewError(ErrNoObjectName, i, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewError(ErrNoObjectDescription, i, v))
		}

		// The default zone does not require exits. New players will spawn here.
		// If it's the only zone in the game, this is perfectly fine. Any other
		// zones without exits are isolated from the game. This might be done
		// deliberately. For instance, a jail to which a game master can
		// teleport players.
		if this.DefaultZone != i && len(v.Exits) == 0 {
			addError(errlist, NewError(ErrZoneIsolated, i, v))
		}
	}

	// Construct the full map grid and make sure each area only has 1 zone defined for it.
	worldmap := NewMap()
	if err := worldmap.Fill(this.Zones); err != nil {
		addError(errlist, err)
	}
}

func (this *World) sanitizeCharacters(errlist *[]*Error) {
	if len(this.Characters) == 0 {
		addError(errlist, NewError(ErrNoCharacters, 0, this))
		return
	}

	for i, v := range this.Characters {
		if len(v.Name) == 0 {
			addError(errlist, NewError(ErrNoObjectName, i, v))
		}

		if v.Class == -1 {
			addError(errlist, NewError(ErrNoCharacterClass, i, v))
		}

		if v.Race == -1 {
			addError(errlist, NewError(ErrNoCharacterRace, i, v))
		}

		if v.Zone == -1 {
			addError(errlist, NewError(ErrCharacterNotPlaced, i, v))
		}
	}
}

func (this *World) sanitizeRaces(errlist *[]*Error) {
	if len(this.Races) == 0 {
		addError(errlist, NewError(ErrNoRaces, 0, this))
		return
	}

	for i, v := range this.Races {
		if len(v.Name) == 0 {
			addError(errlist, NewError(ErrNoObjectName, i, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewError(ErrNoObjectDescription, i, v))
		}
	}
}

func (this *World) sanitizeClasses(errlist *[]*Error) {
	if len(this.Classes) == 0 {
		addError(errlist, NewError(ErrNoClasses, 0, this))
		return
	}

	for i, v := range this.Classes {
		if len(v.Name) == 0 {
			addError(errlist, NewError(ErrNoObjectName, i, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewError(ErrNoObjectDescription, i, v))
		}
	}
}

func (this *World) sanitizeCurrency(errlist *[]*Error) {
	if len(this.Currency) == 0 {
		addError(errlist, NewError(ErrNoCurrency, 0, this))
		return
	}

	for i, v := range this.Currency {
		if len(v.Name) == 0 {
			addError(errlist, NewError(ErrNoObjectName, i, v))
		}

		if v.Value == 0 {
			addError(errlist, NewError(ErrNoCurrencyValue, i, v))
		}

		for j, v1 := range this.Currency {
			if i != j && v.Value == v1.Value {
				addError(errlist, NewError(ErrDuplicateCurrencyValue, i, v))
			}
		}
	}
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
