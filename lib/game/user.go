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
	return v
}

func (this *User) Dispose() {
	if this.Data != nil {
		if this.Info != nil {
			this.Data.SetUser(this.Info)
		}

		this.Data.Close()
		this.Data = nil
	}

	this.Info = nil
}

func (this *User) Login(name, pass string) (err os.Error) {
	if this.Info, err = this.Data.GetUserByName(name); err != nil {
		return
	}

	if this.Info.Password != pass {
		return ErrInvalidCredentials
	}

	return
}

func (this *User) Logout(name, pass string) (err os.Error) {
	if err = this.Data.SetUser(this.Info); err != nil {
		return
	}

	this.Info = nil
	return
}
