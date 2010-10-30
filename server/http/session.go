package main

type Session struct {
	Id    string
	Style string
}

func NewSession(id string) *Session {
	s := new(Session)
	s.Id = id
	s.Style = "default"
	return s
}
