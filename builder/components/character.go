package builder

type Character struct{
	Name        string
	Description string
	Title       string
	Group       *Group
	Class       *Class
	Race        *Race
}
