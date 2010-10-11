package main

import "os"
import "os/signal"
import "mudkip/lib"
import "fmt"

// Any version above this constitutes an incompatible server. This client is
// written to deal with the version 1 API.
const MaxServerVersion = 1

func main() {
	cfg := parseArgs()

	var err os.Error

	client := NewClient(cfg.Secure)
	if err = client.Open(cfg.Server); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
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

func parseArgs() (cfg *Config) {
	var err os.Error
	var cfgfile string

	cfg = NewConfig()

	if len(os.Args) > 1 {
		cfgfile = os.Args[1]
	} else {
		usage()
		os.Exit(0)
	}

	if err = cfg.Load(cfgfile); err != nil {
		fmt.Fprintf(os.Stdout, "Saving template configuration at: %s\n", cfgfile)
		fmt.Fprint(os.Stdout, "Modify it in a text editor and restart this program.\n")

		if err = cfg.Save(cfgfile); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	return
}

func usage() {
	fmt.Fprintf(os.Stdout, `usage: %s <configfile>

   configfile: Full path to a valid configuration profile. If the file does not
               yet exist, the client will create a default template for you in
               the specified location.
`,
		os.Args[0])
}
