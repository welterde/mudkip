package lib

type Currency struct {
	Id    int64
	Name  string
	Value int
}

func NewCurrency() *Currency {
	return new(Currency)
}
