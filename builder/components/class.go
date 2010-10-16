package builder

type Class struct{}

func NewClass() *Class {
	v := new(Class)
	return v
}
