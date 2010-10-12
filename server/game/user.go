package main

import "mudkip/lib"
import "mudkip/store"

type User struct {
	Id   uint16
	Data lib.DataStore
}

func NewUser(dsparams map[string]string) *User {
	v := new(User)
	v.Data = store.New()
	v.Data.Open(dsparams)
	return v
}

func (this *User) Dispose() {
	if this.Data != nil {
		this.Data.Close()
		this.Data = nil
	}
}
