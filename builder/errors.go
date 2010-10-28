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
	ErrCharacterNotPlaced     = "Character has no zone assigned"
)

type Error struct {
	Message string
	Type    string
	Id      int
}

func NewError(m string, id int, obj interface{}) *Error {
	err := new(Error)
	err.Message = m
	err.Id = id
	err.Type = fmt.Sprintf("%T", obj)
	err.Type = err.Type[strings.Index(err.Type, ".")+1:]
	return err
}

func (this *Error) String() string {
	return fmt.Sprintf("%s[%d]: %s", this.Type, this.Id, this.Message)
}

func addError(l *[]*Error, err *Error) {
	sz := len(*l)

	if sz >= cap(*l) {
		cp := make([]*Error, sz, sz+10)
		copy(cp, *l)
		*l = cp
	}

	*l = (*l)[0 : sz+1]
	(*l)[sz] = err
}
