package builder

type Currency struct {
	Name  string
	Value int
}

func NewCurrency() *Currency {
	return new(Currency)
}
