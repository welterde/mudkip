package main

import "os"
import "fmt"
import "os/signal"

var methods *ServiceMethodList
var context *Context

func main() {
	cfg := getConfig()
	context = NewContext(cfg)
	methods = NewServiceMethodList()

	if err := BindApi(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}

	go Run(cfg)

loop:
	for {
		select {
		case sig := <-signal.Incoming:
			if unix, ok := sig.(signal.UnixSignal); ok {
				switch unix {
				case signal.SIGINT, signal.SIGTERM, signal.SIGKILL:
					break loop
				}
			}
		}
	}

	os.Exit(0)
}

func getConfig() (cfg *Config) {
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
		fmt.Fprintf(os.Stdout, "[i] Saving template configuration at: %s\n", cfgfile)
		fmt.Fprint(os.Stdout, "[i] Modify it in a text editor and restart this program.\n")

		if err = cfg.Save(cfgfile); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %v\n", err)
			os.Exit(1)
		}

		os.Exit(0)
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
