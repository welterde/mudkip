package lib

import "os"

type UserInfo struct {
	Id         uint16
	Name       string
	Password   string
	Zone       uint16
	Registered int64
}

type User struct {
	Info *UserInfo
	Data DataStore
}

func NewUser(ds DataStore) *User {
	v := new(User)
	v.Data = ds
	v.Info = new(UserInfo)
	return v
}

func (this *User) Dispose() {
	if this.Data != nil {
		this.Data.Close()
		this.Data = nil
	}
}

func (this *User) Login(name, pass string) (err os.Error) {

	return
}

func (this *User) Logout(name, pass string) (err os.Error) {
	return
}
