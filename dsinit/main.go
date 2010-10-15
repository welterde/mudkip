package main

import "fmt"
import "os"
import "optarg"
import "mudkip/builder"
import "mudkip/store"

func main() {
	cfg := getConfig()

	var err os.Error
	var world *builder.World

	if world, err = builder.LoadWorld(cfg.WorldFile); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}

	ds := store.New()
	if err = ds.Open(cfg.Datastore); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}

	defer ds.Close()

	if err = ds.Initialize(world); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func getConfig() (cfg *Config) {
	var cfgfile string

	cfg = NewConfig()

	optarg.Add("c", "config", "Full path to a configuration file.", "")
	optarg.Add("w", "world", "Full path the world file to process.", "")
	optarg.Add("h", "help", "Displays this help.", "")

	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "c":
			cfgfile = opt.String()
		case "w":
			cfg.WorldFile = opt.String()
		case "h":
			optarg.Usage()
			os.Exit(0)
		}
	}

	if len(cfgfile) == 0 {
		optarg.Usage()
		os.Exit(1)
	}

	if len(cfg.WorldFile) == 0 {
		optarg.Usage()
		os.Exit(1)
	}

	if err := cfg.Load(cfgfile); err != nil {
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
