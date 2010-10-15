package store

import "os"
import "mudkip/lib"

type Store struct{}

// This function name + signature is mandatory. We call it from the server to
// create a new Datastore instance.
func New() lib.DataStore { return new(Store) }

func (this *Store) Open(params map[string]string) (err os.Error) { return }
func (this *Store) Close()                                       {}
func (this *Store) Initialize() (err os.Error)                   { return }

func (this *Store) GetWorldInfo() (info lib.WorldInfo, err os.Error) { return }
func (this *Store) SetWorldInfo(info *lib.WorldInfo) (err os.Error)  { return }

func (this *Store) GetObjectsByType(objtype uint8) (list []lib.Object, err os.Error)  { return }
func (this *Store) GetObject(id uint16, objtype uint8) (obj lib.Object, err os.Error) { return }
func (this *Store) SetObject(lib.Object) (err os.Error)                               { return }

func (this *Store) GetUser(id uint16) (usr *lib.UserInfo, err os.Error)         { return }
func (this *Store) GetUserByName(name string) (usr *lib.UserInfo, err os.Error) { return }
func (this *Store) SetUser(usr *lib.UserInfo) (err os.Error)                    { return }
func (this *Store) GetUsers() (usr []*lib.UserInfo, err os.Error)               { return }
