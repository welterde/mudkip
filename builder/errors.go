package builder

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
	ErrZoneIsolated           = "Zone has no exits"
	ErrNoCurrencyValue        = "Currency value is 0"
	ErrDuplicateCurrencyValue = "Multiple currencies with the same value"
	ErrNoCharacterClass       = "Character has no class"
	ErrNoCharacterRace        = "Character has no race"
)

type Error struct {
	Message string
	Type    string
	Id      int
}

func (this *Error) String() string {
	return fmt.Sprintf("%s[%d]: %s", this.Type, this.Id, this.Message)
}

func addError(l *[]*Error, m string, id int, obj interface{}) {
	sz := len(*l)

	if sz >= cap(*l) {
		cp := make([]*Error, sz, sz+10)
		copy(cp, *l)
		*l = cp
	}

	err := new(Error)
	err.Message = m
	err.Type = fmt.Sprintf("%T", obj)
	err.Id = id
	if idx := strings.Index(err.Type, "."); idx != -1 {
		err.Type = err.Type[idx+1:]
	}

	*l = (*l)[0 : sz+1]
	(*l)[sz] = err
}
