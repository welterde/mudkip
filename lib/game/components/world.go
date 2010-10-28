package lib

type World struct {
	Id            int64
	Created       int64
	Name          string
	Description   string
	Logo          string
	Motd          string
	AllowRegister bool
	LevelCap      int
	DefaultZone   int64
	Zones         ZoneList
	Characters    CharacterList
	Groups        GroupList
	Races         RaceList
	Classes       ClassList
	Currency      CurrencyList
	Armor         ArmorList
	Weapons       WeaponList
	Consumables   ConsumableList
}

func NewWorld() *World {
	v := new(World)
	v.Id = 1
	v.Zones = NewZoneList()
	v.Characters = NewCharacterList()
	v.Races = NewRaceList()
	v.Classes = NewClassList()
	v.Currency = NewCurrencyList()
	v.Groups = NewGroupList()
	v.Armor = NewArmorList()
	v.Weapons = NewWeaponList()
	v.Consumables = NewConsumableList()
	v.AllowRegister = true
	v.DefaultZone = -1
	return v
}
