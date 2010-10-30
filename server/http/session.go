package main

type Session struct {
	Id string
}

func NewSession(id string) *Session {
	s := new(Session)
	s.Id = id
	return s
}
