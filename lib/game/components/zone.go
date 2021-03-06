package lib

/*
A single zone can be anything from a room to a patch of land or space.
It is represented as a square area in which a player can find him/her self
All events/characters/actions take place in discrete zones.

Every zone defines a maximum of 8 exit points, which link to adjascent zones.
By linking one zone to another through these 'portals', one can create an 
arbitrarily large world with different area types and places to visit.

The 8 exit points are all marked by compass directions and a single zone id to
which the portal links.

  NW----N----NE NW----N----NE
  |           | |           |
  |  ZONE  1  | |  ZONE  2  |
  W           E W           E
  |           | |           |
  |           | |           |
  SW----S----SE SW----S----SE
  NW----N----NE NW----N----NE
  |           | |           |
  |  ZONE  3  | |  ZONE  4  |
  W           E W           E
  |           | |           |
  |           | |           |
  SW----S----SE SW----S----SE

One zone's Northern exit links to another zone's Southern exit, etc.

The diagonal exits would technically link to 3 adjascent zones, but to simplify
the navigation in a text-based environment, we opt to strictly use them for
the diagonal route. So from the example above:

 [Zone 1: South-East] <-> [Zone 4: North-West]
 [Zone 2: South-West] <-> [Zone 3: North-East]
 [Zone 3: East      ] <-> [Zone 4: West      ]
 [Zone 2: South     ] <-> [Zone 4: North     ]
 etc.

When a game world has been built, the World.Sanitize() method will traverse all
zones and ensure that all links match up properly. Specifically to make sure we
do not have multiple zones which occupy the same space.

For example, we have Zones 1, 2 and 3 defined. Then add zone 4 which links
[2:S] to [4:N]. Then define zone 5 which links [3:E] to [5:W]. This would put
both zone 4 and 5 in the same location and would constitute a spacial paradox.
Physics 101 teaches us that such things are proposterous and should they be 
attempted, will most likely end in the entire universe collapsing into a super-
massive cupcake. One can understand this would be a bad deal for all involved.

It should be mentioned though, that the warnings generated by this in
World.Sanitize() are not fatal. So if one choses to ignore these warnings, the
game will behave just the way you have defined it to behave. Just be aware that
this will likely confuse the living daylights out of your players.
*/

type Zone struct {
	Id          int64
	Name        string
	Description string
	Lighting    string
	Smell       string
	Sound       string
	Exits       []*Portal
}

func NewZone() *Zone {
	v := new(Zone)
	v.Exits = make([]*Portal, 0, 8)
	return v
}
