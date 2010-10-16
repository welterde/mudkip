package builder

import "os"

var (
	ErrNoWorldName        = os.NewError("World has no name")
	ErrNoWorldDescription = os.NewError("World has no description")

	ErrNoObjectName        = os.NewError("Object has no name")
	ErrNoObjectDescription = os.NewError("Object has no description")

	ErrNoZones      = os.NewError("World has no zones")
	ErrNoCharacters = os.NewError("World has no characters")
	ErrNoClasses    = os.NewError("World has no classes")
	ErrNoRaces      = os.NewError("World has no races")
	ErrNoCurrency   = os.NewError("World has no currency")
)

func addError(l *[]os.Error, e os.Error) {
	sz := len(*l)

	if sz >= cap(*l) {
		cp := make([]os.Error, sz, sz+10)
		copy(cp, *l)
		*l = cp
	}

	*l = (*l)[0 : sz+1]
	(*l)[sz] = e
}
