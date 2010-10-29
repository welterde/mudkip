package lib

type ItemType uint8

// item types
const (
	TypeArmor ItemType = iota
	TypeWeapon
	TypeConsumable
	TypeQuestItem
	TypeResource
	TypeCurrency
)
