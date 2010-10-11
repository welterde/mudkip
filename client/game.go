package main

import "fmt"
import "os"
import "mudkip/lib"

func handleMessage(client *Client, msg lib.Message) {
	switch tt := msg.(type) {
	case *lib.Error:
		fmt.Fprintf(os.Stderr, "Error: %v\n", lib.IntToErr(tt.Errno))

	case *lib.ServerVersion:
		fmt.Fprintf(os.Stdout, "Mudkip Version %d\n", tt.Version)

		if tt.Version > MaxServerVersion {
			fmt.Fprint(os.Stderr, "This client appears to be outdated. We recommend you update it to reflect the latest server version.\n")
			client.Close()
		}

	case *lib.MaxClientsReached:
		fmt.Fprint(os.Stderr, "Maximum number of clients reached.\n")
		client.Close()
	}
}
