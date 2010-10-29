package lib

import "fmt"
import "strings"

const (
	ErrNoObjectName           = "Object has no name"
	ErrNoObjectDescription    = "Object has no description"
	ErrNoZones                = "World has no zones"
	ErrNoDefaultZone          = "World has no default zone"
	ErrNoCharacters           = "World has no characters"
	ErrNoClasses              = "World has no classes"
	ErrNoRaces                = "World has no races"
	ErrNoCurrency             = "World has no currency"
	ErrNoArmor                = "World has no armor"
	ErrNoWeapons              = "World has no weapons"
	ErrNoConsumables          = "World has no consumables"
	ErrNoQuestItems           = "World has no questitems"
	ErrNoQuests               = "World has no quests"
	ErrZoneIsolated           = "Zone has no exits"
	ErrNoCurrencyValue        = "Currency value is 0"
	ErrDuplicateCurrencyValue = "Multiple currencies with the same value"
	ErrNoCharacterClass       = "Character has no class"
	ErrNoCharacterRace        = "Character has no race"
	ErrCharacterNotPlaced     = "Character has no zone assigned"
	ErrInventoryFull          = "Inventory is full"
	ErrWeaponNoDamage         = "Weapon has no damage range defined"
	ErrQuestItemNoQuest       = "Quest item is not associated with a quest"
	ErrQuestNoRewards         = "Quest has no rewards"
	ErrQuestNoSource          = "Quest is not associated with a character"
)

type BuildError struct {
	Message string
	Type    string
	Id      int64
}

func NewBuildError(m string, id int64, obj interface{}) *BuildError {
	err := new(BuildError)
	err.Message = m
	err.Id = id
	err.Type = fmt.Sprintf("%T", obj)
	err.Type = err.Type[strings.Index(err.Type, ".")+1:]
	return err
}

func (this *BuildError) String() string {
	return fmt.Sprintf("%s[%d]: %s", this.Type, this.Id, this.Message)
}

func addError(l *[]*BuildError, err *BuildError) {
	sz := len(*l)

	if sz >= cap(*l) {
		cp := make([]*BuildError, sz, sz+10)
		copy(cp, *l)
		*l = cp
	}

	*l = (*l)[0 : sz+1]
	(*l)[sz] = err
}
