package store

import "os"
import "mudkip/lib"

type Store struct{}

// This function name + signature is mandatory. We call it from the server to
// create a new Datastore instance. If need be, any additional initialization
// bits can be put in here.
func New() lib.DataStore { return new(Store) }

func (this *Store) Open(map[string]string) os.Error                 { return nil }
func (this *Store) Close()                                          {}
func (this *Store) Initialize(*lib.World) os.Error                  { return nil }
func (this *Store) GetArmor(int64) (*lib.Armor, os.Error)           { return nil, nil }
func (this *Store) SetArmor(*lib.Armor) os.Error                    { return nil }
func (this *Store) GetCharacter(int64) (*lib.Character, os.Error)   { return nil, nil }
func (this *Store) SetCharacter(*lib.Character) os.Error            { return nil }
func (this *Store) GetClass(int64) (*lib.Class, os.Error)           { return nil, nil }
func (this *Store) SetClass(*lib.Class) os.Error                    { return nil }
func (this *Store) GetConsumable(int64) (*lib.Consumable, os.Error) { return nil, nil }
func (this *Store) SetConsumable(*lib.Consumable) os.Error          { return nil }
func (this *Store) GetCurrency(int64) (*lib.Currency, os.Error)     { return nil, nil }
func (this *Store) SetCurrency(*lib.Currency) os.Error              { return nil }
func (this *Store) GetGroup(int64) (*lib.Group, os.Error)           { return nil, nil }
func (this *Store) SetGroup(*lib.Group) os.Error                    { return nil }
func (this *Store) GetInventory(int64) (*lib.Inventory, os.Error)   { return nil, nil }
func (this *Store) SetInventory(*lib.Inventory) os.Error            { return nil }
func (this *Store) GetPortal(int64) (*lib.Portal, os.Error)         { return nil, nil }
func (this *Store) SetPortal(*lib.Portal) os.Error                  { return nil }
func (this *Store) GetRace(int64) (*lib.Race, os.Error)             { return nil, nil }
func (this *Store) SetRace(*lib.Race) os.Error                      { return nil }
func (this *Store) GetStats(int64) (*lib.Stats, os.Error)           { return nil, nil }
func (this *Store) SetStats(*lib.Stats) os.Error                    { return nil }
func (this *Store) GetWeapon(int64) (*lib.Weapon, os.Error)         { return nil, nil }
func (this *Store) SetWeapon(*lib.Weapon) os.Error                  { return nil }
func (this *Store) GetWorld() (*lib.World, os.Error)                { return nil, nil }
func (this *Store) SetWorld(*lib.World) os.Error                    { return nil }
func (this *Store) GetZone(int64) (*lib.Zone, os.Error)             { return nil, nil }
func (this *Store) SetZone(*lib.Zone) os.Error                      { return nil }
func (this *Store) GetUser(int64) (*lib.UserInfo, os.Error)         { return nil, nil }
func (this *Store) GetUserByName(string) (*lib.UserInfo, os.Error)  { return nil, nil }
func (this *Store) SetUser(*lib.UserInfo) os.Error                  { return nil }
func (this *Store) GetUsers() ([]*lib.UserInfo, os.Error)           { return nil, nil }
