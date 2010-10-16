package builder

type Currency struct {
	Name  string
	Value int
}

func NewCurrency(name string, val int) *Currency {
	v := new(Currency)
	v.Name = name
	v.Value = val
	return v
}
