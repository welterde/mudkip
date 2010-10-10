package client

import "testing"
import "os"
import "os/signal"
import "mudkip/lib"
import "fmt"

// Any version above this constitutes an incompatible server. This client is
// written to deal with the version 1 API.
const MaxServerVersion = 1

func Test(t *testing.T) {
	var err os.Error

	client := NewClient()
	if err = client.Open("0.0.0.0:54321"); err != nil {
		t.Error(err.String())
		return
	}

	defer client.Close()

	var msg lib.Message
	var sig signal.Signal

loop:
	for {
		select {
		case msg = <-client.Messages:
			go handleMessage(client, msg)

		case sig = <-signal.Incoming:
			if unix, ok := sig.(signal.UnixSignal); ok {
				switch unix {
				case signal.SIGINT, signal.SIGTERM, signal.SIGKILL:
					return
				}
			}
		}

		if closed(client.Messages) || closed(signal.Incoming) {
			return
		}
	}
}

func handleMessage(client *Client, msg lib.Message) {
	switch tt := msg.(type) {
	case *lib.ServerVersion:
		fmt.Printf("Mudkip Version %d\n", tt.Version)

		if tt.Version > MaxServerVersion {
			client.Close() // Unsupported.
		}

	case *lib.MaxClientsReached:
		fmt.Print("Maximum number of clients reached.\n")
		client.Close()
	}
}
