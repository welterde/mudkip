package main

import "strings"

type HttpMethod uint8

const (
	GET HttpMethod = iota
	POST
	HEAD
	CONNECT
	DELETE
	OPTIONS
	PUT
	TRACE
)

func (this HttpMethod) Equals(name string) bool {
	switch strings.ToUpper(name) {
	case "GET":
		return this == GET
	case "POST":
		return this == POST
	case "HEAD":
		return this == HEAD
	case "CONNECT":
		return this == CONNECT
	case "DELETE":
		return this == DELETE
	case "OPTIONS":
		return this == OPTIONS
	case "PUT":
		return this == PUT
	case "TRACE":
		return this == TRACE
	}
	return false
}
