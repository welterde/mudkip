package lib

import "testing"
import "os"

func TestBuild(t *testing.T) {
	world := NewWorld()
	world.Name = "Mudkipia"
	world.Description = "Magical land of unicorns and mudkipz"
	world.LevelCap = 20
	world.AllowRegister = true
	world.Logo = `
    ROFL:ROFL:LOL:ROFL:ROFL
         ______|_____
 L      /            \
LOL=====            []\
 L      \______________\
           |     |
         ============/
`

	zone := NewZone()
	zone.Name = "Seaside Tavern"
	zone.Description = "Welcome area for tourists and would be adventurers. Feel free to hang around and chat.\n\nWhen you are ready to start hunting for the elusive unicorn, please speak to the proprietor in the corner to get your hunting license sorted."
	zone.Lighting = "dark, candle lit with a shimmering open fire in the northern corner"
	zone.Smell = "of beer and stale bread"
	zone.Sound = "quiet whispers and muttering"
	world.DefaultZone = world.Zones.Add(zone)

	class := NewClass()
	class.Name = "Warrior"
	class.Description = "A battle hardened fighter"
	class.StatBonus.HP = 100
	class.StatBonus.AP = 30
	class.StatBonus.DEF = 40
	class.StatBonus.STR = 40
	warrior := world.Classes.Add(class)

	race := NewRace()
	race.Name = "Human"
	race.Description = "Versatile, greedy and generally obnoxious"
	race.StatBonus.HP = 10
	race.StatBonus.MP = 10
	race.StatBonus.AP = 5
	race.StatBonus.DEF = 5
	race.StatBonus.AGI = 5
	race.StatBonus.STR = 5
	race.StatBonus.WIS = 5
	race.StatBonus.LUC = 5
	race.StatBonus.CHR = 5
	race.StatBonus.PER = 5
	human := world.Races.Add(race)

	char := NewCharacter()
	char.Name = "bob"
	char.Title = "Bringer of Doom"
	char.Level = 1
	char.Class = warrior
	char.Race = human
	char.Zone = world.DefaultZone
	world.Characters.Add(char)

	copper := NewCurrency()
	copper.Name = "copper"
	copper.Value = 1
	world.Currency.Add(copper)

	silver := NewCurrency()
	silver.Name = "silver"
	silver.Value = 100 * copper.Value
	world.Currency.Add(silver)

	gold := NewCurrency()
	gold.Name = "gold"
	gold.Value = 100 * silver.Value
	world.Currency.Add(gold)

	weapon := NewWeapon()
	weapon.Name = "Club of a thousand pains"
	weapon.Description = "This club has seen much anguish. You can see some bodily remains stuck inbetween the cracks in the wood."
	weapon.Damage = [2]int{0, 100}
	weapon.Type = Melee | TwoHanded
	weapon.StatBonus.AP = 10
	weapon.StatBonus.STR = 20
	world.Weapons.Add(weapon)

	armor := NewArmor()
	armor.Name = "Tunic of the smelly vagrant"
	armor.Description = "You should really wash this before wearing it. Who knows where it has been..."
	armor.Type = Chest
	armor.StatBonus.HP = 10
	armor.StatBonus.CHR = -10
	world.Armor.Add(armor)

	bread := NewConsumable()
	bread.Liquid = false
	bread.Name = "Bread"
	bread.Description = "An old fashioned, hearty and nutritional loaf of bread"
	bread.StatBonus.HP = 10
	world.Consumables.Add(bread)

	tea := NewConsumable()
	tea.Liquid = true
	tea.Name = "Ginger Tea"
	tea.Description = "A nice, hot cup of tea"
	tea.StatBonus.MP = 10
	world.Consumables.Add(tea)

	if errlist := SanitizeWorld(world); len(errlist) > 0 {
		for _, err := range errlist {
			t.Errorf("%v", err)
		}
		return
	}

	var err os.Error
	if err = SaveWorld("test.js", world, false); err != nil {
		t.Error(err.String())
	}

	if world, err = LoadWorld("test.js"); err != nil {
		t.Error(err.String())
	}
}
