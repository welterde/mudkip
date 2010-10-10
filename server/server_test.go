package server

import "testing"
import "os"
import "os/signal"
import "mudkip/lib"
import "fmt"

func Test(t *testing.T) {
	var err os.Error

	srv := NewServer(os.Stdout, false, 16)

	if err = srv.Open("0.0.0.0:54321"); err != nil {
		t.Error(err.String())
		return
	}

	var msg lib.Message
	var sig signal.Signal

loop:
	for {
		select {
		case msg = <-srv.Messages:
			switch tt := msg.(type) {
			case *lib.ClientConnected:
				fmt.Printf("Client connected: %s\n", msg.Sender())

			case *lib.ClientDisconnected:
				fmt.Printf("Client disconnected: %s\n", msg.Sender())
			}

		case sig = <-signal.Incoming:
			if unix, ok := sig.(signal.UnixSignal); ok {
				switch unix {
				case signal.SIGINT, signal.SIGTERM, signal.SIGKILL:
					break loop
				}
			}
		}

		if closed(srv.Messages) || closed(signal.Incoming) {
			break loop
		}
	}

	srv.Close()
}
