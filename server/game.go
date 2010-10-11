package main

import "mudkip/lib"

func handleMessage(srv *Server, msg lib.Message) {
	switch tt := msg.(type) {
	case *lib.ClientConnected:
		srv.Info("Client connected: %s", msg.Sender())

	case *lib.ClientDisconnected:
		srv.Info("Client disconnected: %s", msg.Sender())
	}
}
