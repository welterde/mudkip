package lib

import "os"

// This interface is the front-end for a compile-time selectable datastore
// backend. This would be a database or binary data file of some sort. The
// actual implementation does the interaction with a given datastore.
// For instance a MysqlStore would implement the methods of this interface to
// let us persist all the game data to an existing mysql database.
type DataStore interface {
	Open(map[string]string) os.Error
	Close()

	Initialize(*World) os.Error

	GetArmor(int64) (*Armor, os.Error)
	SetArmor(*Armor) os.Error

	GetCharacter(int64) (*Character, os.Error)
	SetCharacter(*Character) os.Error

	GetClass(int64) (*Class, os.Error)
	SetClass(*Class) os.Error

	GetConsumable(int64) (*Consumable, os.Error)
	SetConsumable(*Consumable) os.Error

	GetCurrency(int64) (*Currency, os.Error)
	SetCurrency(*Currency) os.Error

	GetGroup(int64) (*Group, os.Error)
	SetGroup(*Group) os.Error

	GetInventory(int64) (*Inventory, os.Error)
	SetInventory(*Inventory) os.Error

	GetPortal(int64) (*Portal, os.Error)
	SetPortal(*Portal) os.Error

	GetQuest(int64) (*Quest, os.Error)
	SetQuest(*Quest) os.Error

	GetQuestItem(int64) (*QuestItem, os.Error)
	SetQuestItem(*QuestItem) os.Error

	GetQuestReward(int64) (*QuestReward, os.Error)
	SetQuestReward(*QuestReward) os.Error

	GetRace(int64) (*Race, os.Error)
	SetRace(*Race) os.Error

	GetStats(int64) (*Stats, os.Error)
	SetStats(*Stats) os.Error

	GetWeapon(int64) (*Weapon, os.Error)
	SetWeapon(*Weapon) os.Error

	GetWorld() (*World, os.Error)
	SetWorld(*World) os.Error

	GetZone(int64) (*Zone, os.Error)
	SetZone(*Zone) os.Error

	GetUser(int64) (*UserInfo, os.Error)
	GetUserByName(string) (*UserInfo, os.Error)
	SetUser(*UserInfo) os.Error
	GetUsers() ([]*UserInfo, os.Error)
}
