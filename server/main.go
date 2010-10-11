package main

import "os"
import "os/signal"
import "mudkip/lib"
import "fmt"

func main() {
	var err os.Error

	cfg, ds := getConfig()
	defer ds.Close()

	srv := NewServer(cfg)
	if err = srv.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	srv.Info("Listening on: %s", srv.conn.Addr())
	srv.Info("Max clients: %d", cfg.MaxClients)
	srv.Info("Client timeout: %d minute(s)", cfg.ClientTimeout)
	srv.Info("Secure connection: %v", cfg.Secure)
	srv.Info("Using datastore: %T", lib.GetStore())

	var msg lib.Message
	var sig signal.Signal

	incoming := srv.Messages()

loop:
	for {
		select {
		case msg = <-incoming:
			go handleMessage(srv, msg)

		case sig = <-signal.Incoming:
			if unix, ok := sig.(signal.UnixSignal); ok {
				switch unix {
				case signal.SIGINT, signal.SIGTERM, signal.SIGKILL:
					break loop
				}
			}
		}

		if closed(incoming) || closed(signal.Incoming) {
			break loop
		}
	}

	srv.Close()
}

func getConfig() (cfg *Config, ds lib.DataStore) {
	var err os.Error
	var cfgfile string

	if len(os.Args) > 1 {
		cfgfile = os.Args[1]
	} else {
		usage()
		os.Exit(0)
	}

	cfg = NewConfig()
	if err = cfg.Load(cfgfile); err != nil {
		fmt.Fprintf(os.Stdout, "Saving template configuration at: %s\n", cfgfile)
		fmt.Fprint(os.Stdout, "Modify it in a text editor and restart this program.\n")

		if err = cfg.Save(cfgfile); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	if len(cfg.Datastore) == 0 {
		fmt.Fprint(os.Stderr, "Missing datastore driver name in config file.\n")
		os.Exit(1)
	}

	if ds = lib.GetStore(); ds == nil {
		fmt.Fprintf(os.Stderr, "Server has been built without datastore support. Cannot continue.\n")
		os.Exit(1)
	}

	if err = ds.Open(cfg.Datastore); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return
}

func usage() {
	fmt.Fprintf(os.Stdout, `usage: %s <configfile>

   configfile: Full path to a valid configuration profile. If the file does not
               yet exist, the server will create a default template for you in
               the specified location.
`,
		os.Args[0])
}
