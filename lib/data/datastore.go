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

	// This method is only used when we initialize a new game world. It should
	// take care of doing all the initial table creation and initialization,
	// relevant to the implemented datastore. What to implement can be seen in
	// the mudkip/data/STRUCTURE file. It shows an extensive overview of the
	// mudkip game structure. You should be able to extrapolate from this, how
	// best to build your data model. We do not call this in the server itself,
	// but from a world builder tool which allows a user to create a new world
	// from scratch.
	Initialize() os.Error

	// Fetches the configuration settings for this game's world. Should be only
	// one of these per datastore.
	GetWorldInfo() (*WorldInfo, os.Error)

	// Add or update the world configuration settings.
	SetWorldInfo(*WorldInfo) os.Error

	// This is a generic function which fetches an object from the datastore
	// and transforms it into the appropriate unpacked type.
	GetObject(uint16, uint8) (Object, os.Error)

	// Selects all objects of the given type
	GetObjectsByType(uint8) ([]Object, os.Error)

	// This stores the given object in the datastore
	SetObject(Object) os.Error

	// Fetches the user info associated with the given id
	GetUser(uint16) (*UserInfo, os.Error)

	// Fetches the user info associated with the given name
	GetUserByName(string) (*UserInfo, os.Error)

	// Stores the given user info
	SetUser(*UserInfo) os.Error

	// Lists all users
	GetUsers() ([]*UserInfo, os.Error)
}
