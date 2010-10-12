package main

import "os"
import "mudkip/lib"
import "mudkip/store"

type User struct {
	PeerId string
	DbId   uint16
	Data   lib.DataStore
}

func NewUser(peerid string, dsparams map[string]string) (v *User, err os.Error) {
	v = new(User)
	v.PeerId = peerid
	v.Data = store.New()
	return v, v.Data.Open(dsparams)
}

func (this *User) Dispose() {
	if this.Data != nil {
		this.Data.Close()
		this.Data = nil
	}
}
