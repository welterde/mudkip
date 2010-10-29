package lib

import "os"
import "path"
import "json"
import "io/ioutil"

// Load a world from a JSON formatted data file.
func LoadWorld(file string) (world *World, err os.Error) {
	var data []byte

	if data, err = ioutil.ReadFile(path.Clean(file)); err != nil {
		return
	}

	world = new(World)
	err = json.Unmarshal(data, &world)
	return
}

// Saves the world to a JSON formatted data file. Optionally with indentation
// for easy reading/modification.
func SaveWorld(file string, world *World, compact bool) (err os.Error) {
	var data []byte

	if compact {
		if data, err = json.Marshal(world); err != nil {
			return
		}
	} else {
		if data, err = json.MarshalIndent(world, "", "  "); err != nil {
			return
		}
	}

	return ioutil.WriteFile(path.Clean(file), data, 0600)
}


// This function goes through the entire data structure and finds irregularities
// in any of the components. Duplicate objects, unlinked zones, inconsistent
// bits and bobs etc and reports them as a list of builder.Error. These are not
// necessarily fatal errors. This will depend on the nature of the game being
// implemented and the wishes of the game master. These errors should just be
// considered a guide to the correct formation of a game world.
func SanitizeWorld(w *World) (errlist []*BuildError) {
	errlist = make([]*BuildError, 0, 10)

	if len(w.Name) == 0 {
		addError(&errlist, NewBuildError(ErrNoObjectName, 0, w))
	}

	if len(w.Description) == 0 {
		addError(&errlist, NewBuildError(ErrNoObjectDescription, 0, w))
	}

	if w.DefaultZone <= 0 {
		addError(&errlist, NewBuildError(ErrNoDefaultZone, 0, w))
	}

	sanitizeZones(w, &errlist)
	sanitizeCharacters(w, &errlist)
	sanitizeRaces(w, &errlist)
	sanitizeClasses(w, &errlist)
	sanitizeCurrency(w, &errlist)
	sanitizeArmor(w, &errlist)
	sanitizeWeapons(w, &errlist)
	sanitizeConsumables(w, &errlist)
	sanitizeQuestItems(w, &errlist)
	sanitizeQuests(w, &errlist)
	return
}

func sanitizeZones(w *World, errlist *[]*BuildError) {
	if len(w.Zones) == 0 {
		addError(errlist, NewBuildError(ErrNoZones, 0, w))
		return
	}

	for _, v := range w.Zones {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectDescription, v.Id, v))
		}

		// The default zone does not require exits. New players will spawn here.
		// If it's the only zone in the game, this is perfectly fine. Any other
		// zones without exits are isolated from the game. This might be done
		// deliberately. For instance, a jail to which a game master can
		// teleport players.
		if w.DefaultZone != v.Id && len(v.Exits) == 0 {
			addError(errlist, NewBuildError(ErrZoneIsolated, v.Id, v))
		}
	}

	// Construct the full map grid and make sure each area only has 1 zone defined for it.
	worldmap := NewMap()
	if err := worldmap.Fill(w.Zones); err != nil {
		addError(errlist, NewBuildError(err.String(), 0, w))
	}
}

func sanitizeCharacters(w *World, errlist *[]*BuildError) {
	if len(w.Characters) == 0 {
		addError(errlist, NewBuildError(ErrNoCharacters, 0, w))
		return
	}

	for _, v := range w.Characters {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if v.Class == -1 {
			addError(errlist, NewBuildError(ErrNoCharacterClass, v.Id, v))
		}

		if v.Race == -1 {
			addError(errlist, NewBuildError(ErrNoCharacterRace, v.Id, v))
		}

		if v.Zone == -1 {
			addError(errlist, NewBuildError(ErrCharacterNotPlaced, v.Id, v))
		}
	}
}

func sanitizeRaces(w *World, errlist *[]*BuildError) {
	if len(w.Races) == 0 {
		addError(errlist, NewBuildError(ErrNoRaces, 0, w))
		return
	}

	for _, v := range w.Races {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectDescription, v.Id, v))
		}
	}
}

func sanitizeClasses(w *World, errlist *[]*BuildError) {
	if len(w.Classes) == 0 {
		addError(errlist, NewBuildError(ErrNoClasses, 0, w))
		return
	}

	for _, v := range w.Classes {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectDescription, v.Id, v))
		}
	}
}

func sanitizeCurrency(w *World, errlist *[]*BuildError) {
	if len(w.Currency) == 0 {
		addError(errlist, NewBuildError(ErrNoCurrency, 0, w))
		return
	}

	var v, v1 *Currency
	for _, v = range w.Currency {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if v.Value == 0 {
			addError(errlist, NewBuildError(ErrNoCurrencyValue, v.Id, v))
		}

		for _, v1 = range w.Currency {
			if v.Id != v1.Id && v.Value == v1.Value {
				addError(errlist, NewBuildError(ErrDuplicateCurrencyValue, v.Id, v))
			}
		}
	}
}

func sanitizeArmor(w *World, errlist *[]*BuildError) {
	if len(w.Armor) == 0 {
		addError(errlist, NewBuildError(ErrNoArmor, 0, w))
		return
	}

	for _, v := range w.Armor {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}
	}
}

func sanitizeWeapons(w *World, errlist *[]*BuildError) {
	if len(w.Weapons) == 0 {
		addError(errlist, NewBuildError(ErrNoWeapons, 0, w))
		return
	}

	for _, v := range w.Weapons {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if v.Damage[0] == 0 && v.Damage[1] == 0 {
			addError(errlist, NewBuildError(ErrWeaponNoDamage, v.Id, v))
		}
	}
}

func sanitizeConsumables(w *World, errlist *[]*BuildError) {
	if len(w.Consumables) == 0 {
		addError(errlist, NewBuildError(ErrNoConsumables, 0, w))
		return
	}

	for _, v := range w.Consumables {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}
	}
}

func sanitizeQuestItems(w *World, errlist *[]*BuildError) {
	if len(w.QuestItems) == 0 {
		addError(errlist, NewBuildError(ErrNoQuestItems, 0, w))
		return
	}

	for _, v := range w.QuestItems {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}
	}
}

func sanitizeQuests(w *World, errlist *[]*BuildError) {
	if len(w.Quests) == 0 {
		addError(errlist, NewBuildError(ErrNoQuests, 0, w))
		return
	}

	for _, v := range w.Quests {
		if len(v.Name) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if len(v.Description) == 0 {
			addError(errlist, NewBuildError(ErrNoObjectName, v.Id, v))
		}

		if v.Character <= 0 {
			addError(errlist, NewBuildError(ErrQuestNoSource, v.Id, v))
		}

		if v.Rewards.Len() == 0 {
			addError(errlist, NewBuildError(ErrQuestNoRewards, v.Id, v))
		}
	}
}
