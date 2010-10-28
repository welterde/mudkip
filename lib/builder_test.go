package lib

import "testing"

func Test(t *testing.T) {
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
	world.DefaultZone = world.AddZone(zone)

	class := NewClass()
	class.Name = "Warrior"
	class.Description = "A battle hardened fighter"
	warrior := world.AddClass(class)

	race := NewRace()
	race.Name = "Human"
	race.Description = "Versatile, greedy and generally obnoxious"
	human := world.AddRace(race)

	char := NewCharacter()
	char.Name = "bob"
	char.Title = "Bringer of Doom"
	char.Level = 1
	char.Class = warrior
	char.Race = human
	char.Zone = world.DefaultZone
	world.AddCharacter(char)

	copper := NewCurrency()
	copper.Name = "copper"
	copper.Value = 1
	world.AddCurrency(copper)

	silver := NewCurrency()
	silver.Name = "silver"
	silver.Value = 100 * copper.Value
	world.AddCurrency(silver)

	gold := NewCurrency()
	gold.Name = "gold"
	gold.Value = 100 * silver.Value
	world.AddCurrency(gold)

	weapon := NewWeapon()
	weapon.Name = "Club of a thousand pains"
	weapon.Description = "This club has seen much anguish. You can see some bodily remains stuck inbetween the cracks in the wood."
	weapon.Damage = [2]int{0, 100}
	weapon.Type = Melee | TwoHanded
	world.AddWeapon(weapon)

	armor := NewArmor()
	armor.Name = "Tunic of the smelly vagrant"
	armor.Description = "You should really wash this before wearing it. Who knows where it has been..."
	armor.Type = Chest
	world.AddArmor(armor)

	if errlist := world.Sanitize(); len(errlist) > 0 {
		for _, err := range errlist {
			t.Errorf("%v", err)
		}
		return
	}

	if err := SaveWorld("test.js", world, false); err != nil {
		t.Error(err.String())
	}
}
