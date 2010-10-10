package server

import "os"

var (
	ErrInvalidClientID = os.NewError("Invalid Clientid specified")
	ErrInvalidPacket   = os.NewError("Invalid packet format")
	ErrMaxClients      = os.NewError("Maximum number of clients reached")
)
