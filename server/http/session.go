package main

type Session struct {
	Id         string
	Style      string
	Registered bool
}

func NewSession(id string) *Session {
	s := new(Session)
	s.Id = id
	s.Style = "default"
	s.Registered = false
	return s
}
