package main

import "os"
import "fmt"
import "os/signal"

var methods *ServiceMethodList
var templates *TemplateCache
var context *ServerContext

func main() {
	config := getConfig()
	context = NewServerContext(config)
	methods = NewServiceMethodList()
	templates = NewTemplateCache()

	if err := templates.Load(config.WebRoot); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}

	if err := BindApi(methods); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}

	go Run(config)

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
	if len(cfg.ListenAddr) == 0 {
		fmt.Fprint(os.Stderr, "[e] No listen address has been specified in the configuration file.\n")
		os.Exit(1)
	}

	if cfg.Secure && (len(cfg.ServerCert) == 0 || len(cfg.ServerKey) == 0) {
		fmt.Fprint(os.Stderr, "[e] When running as a secure server, you must specify valid "+
			"servercert and serverkey values. These should point to files containing the "+
			"respective certificate and key.\n")
		os.Exit(1)
	}

	if len(cfg.CookieSalt) == 0 {
		fmt.Fprint(os.Stderr, "[e] It is highly recommended to set a valid cookiesalt value in "+
			"the configuration file. This salt is used to (de/en)crypt cookies. It should be an "+
			"arbitrary length, random string of characters.\n")
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
