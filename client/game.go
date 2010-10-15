package main

import "fmt"
import "os"
import "mudkip/lib"

// Any version above this constitutes an incompatible server. This client is
// written to deal with the version 1 API.
const MaxServerVersion = 1

func handleMessage(client *Client, msg lib.Message) {
	fmt.Fprintf(os.Stderr, "%s -> %T\n", msg.Sender(), msg)

	switch tt := msg.(type) {
	case *lib.Ok:

	case *lib.EnterZone:

	case *lib.LeaveZone:

	case *lib.Error:
		fmt.Fprintf(os.Stderr, "Error: %v\n", tt.ToError())

	case *lib.ServerVersion:
		if tt.Version > MaxServerVersion {
			fmt.Fprint(os.Stderr, "This client appears to be outdated. We recommend you update it to reflect the latest server version.\n")
			client.Close()
		}

	case *lib.MaxClientsReached:
		fmt.Fprint(os.Stderr, "Maximum number of clients reached.\n")
		client.Close()
	}
}
