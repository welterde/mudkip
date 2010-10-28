package main

import "fmt"
import "os"
import "mudkip/lib"
import "mudkip/store"

func main() {
	context := NewContext(getConfig())
	defer context.Dispose()

	fmt.Fprint(os.Stdout, "[i] Running server...\n\n")

	if err := context.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
	}

	return
}

func getConfig() (cfg *Config, world *lib.World) {
	var err os.Error
	var cfgfile string

	if len(os.Args) > 1 {
		cfgfile = os.Args[1]
	} else {
		usage()
		os.Exit(0)
	}

	fmt.Fprint(os.Stdout, "[i] Reading configuration...\n")

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

	var ds lib.DataStore

	fmt.Fprint(os.Stdout, "[i] Testing datastore...\n")

	// Make sure we have a valid datastore available.
	if ds = store.New(); ds == nil {
		fmt.Fprint(os.Stderr, "[e] Server has been built without valid datastore support. Cannot continue.\n")
		os.Exit(1)
	}

	if err = ds.Open(cfg.Datastore); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}

	defer ds.Close()

	if world, err = ds.GetWorld(); err != nil {
		fmt.Fprint(os.Stderr, "[e] It appears that there is no valid world info available.\n")
		fmt.Fprint(os.Stderr, "[e] Have you initialized the datastore with mudkip/dsinit?\n")
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
