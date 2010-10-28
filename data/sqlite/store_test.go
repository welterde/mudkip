package store

import "testing"
import "mudkip/builder"

func Test(t *testing.T) {
	world := builder.NewWorld()
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

	zone := builder.NewZone()
	zone.Name = "Seaside Tavern"
	zone.Description = "Welcome area for tourists and would be adventurers. Feel free to hang around and chat.\n\nWhen you are ready to start hunting for the elusive unicorn, please speak to the proprietor in the corner to get your hunting license sorted."
	zone.Lighting = "dark, candle lit with a shimmering open fire in the northern corner"
	zone.Smell = "of beer and stale bread"
	zone.Sound = "quiet whispers and muttering"
	world.DefaultZone = world.AddZone(zone)

	class := builder.NewClass()
	class.Name = "Warrior"
	class.Description = "A battle hardened fighter"
	warrior := world.AddClass(class)

	race := builder.NewRace()
	race.Name = "Human"
	race.Description = "Versatile, greedy and generally obnoxious"
	human := world.AddRace(race)

	char := builder.NewCharacter()
	char.Name = "bob"
	char.Title = "Bringer of Doom"
	char.Level = 1
	char.Class = warrior
	char.Race = human
	char.Zone = world.DefaultZone
	world.AddCharacter(char)

	copper := builder.NewCurrency()
	copper.Name = "copper"
	copper.Value = 1
	world.AddCurrency(copper)

	silver := builder.NewCurrency()
	silver.Name = "silver"
	silver.Value = 100 * copper.Value
	world.AddCurrency(silver)

	gold := builder.NewCurrency()
	gold.Name = "gold"
	gold.Value = 100 * silver.Value
	world.AddCurrency(gold)

	if errlist := world.Sanitize(); len(errlist) > 0 {
		var err *builder.Error
		for _, err = range errlist {
			t.Errorf("%v", err)
		}
		return
	}

	store := New()
	store.Open(map[string]string{"file": "test.db"})
	defer store.Close()

	if err := store.Initialize(world); err != nil {
		t.Error(err.String())
	}
}
