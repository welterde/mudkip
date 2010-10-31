package main

type MenuItem struct {
	Url      string
	Title    string
	Name     string
	Selected bool
}

func NewMenuItem(url, title, name string, selected bool) *MenuItem {
	m := new(MenuItem)
	m.Url = url
	m.Title = title
	m.Name = name
	m.Selected = selected
	return m
}

var (
	mainMenu     []*MenuItem
	accountMenuA []*MenuItem
	accountMenuB []*MenuItem
	worldsMenu   []*MenuItem
)

func init() {
	mainMenu = []*MenuItem{
		NewMenuItem("/", "Go to home page", "Home", false),
		NewMenuItem("/worlds", "Browser worlds", "Worlds", false),
		NewMenuItem("/account", "Create or manage your account", "Account", false),
	}

	accountMenuA = []*MenuItem{ // account menu for non-logged in user
		NewMenuItem("/account", "Account overview", "Overview", false),
		NewMenuItem("/account/register", "Register a new account", "Register", false),
		NewMenuItem("/account/login", "Login to your account", "Login", false),
	}

	accountMenuB = []*MenuItem{ // account menu for logged in user
		NewMenuItem("/account", "Account overview", "Overview", false),
		NewMenuItem("/account/edit", "Edit your account", "Edit", false),
		NewMenuItem("/account/logout", "Logout from your account", "Logout", false),
	}

	worldsMenu = []*MenuItem{
		NewMenuItem("/worlds", "World overview", "Overview", false),
		NewMenuItem("/worlds/play", "Play in one of the worlds", "Play", false),
		NewMenuItem("/worlds/create", "Create a new world", "Create", false),
		NewMenuItem("/worlds/edit", "Edit an existing world", "Edit", false),
	}
}
