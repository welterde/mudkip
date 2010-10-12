package main

import "fmt"
import "os"
import "mudkip/lib"
import "mudkip/store"

func main() {
	context := NewContext(getConfig())
	defer context.Dispose()

	if err := context.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	return
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
		fmt.Fprintf(os.Stdout, "Saving template configuration at: %s\n", cfgfile)
		fmt.Fprint(os.Stdout, "Modify it in a text editor and restart this program.\n")

		if err = cfg.Save(cfgfile); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	var ds lib.DataStore

	// Make sure we have a valid datastore available.
	if ds = store.New(); ds == nil {
		fmt.Fprintf(os.Stderr, "Server has been built without datastore support. Cannot continue.\n")
		os.Exit(1)
	}

	if err = ds.Open(cfg.Datastore); err != nil {
		fmt.Fprintf(os.Stderr, "Datastore failure: %v\n", err)
		os.Exit(1)
	}

	ds.Close()

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
