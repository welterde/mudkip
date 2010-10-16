package builder

type Class struct{
	Name        string
	Description string
}

func NewClass() *Class {
	v := new(Class)
	return v
}
