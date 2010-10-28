package main

import "os"
import "ini"

type Config struct {
	WorldFile string
	Datastore map[string]string
}

func NewConfig() *Config {
	c := new(Config)
	c.Datastore = make(map[string]string)
	return c
}

func (this *Config) Load(file string) (err os.Error) {
	var cfg *ini.Config
	var data *ini.Section
	var ok bool

	if cfg, err = ini.Load(file); err != nil {
		return
	}

	if data, ok = cfg.Sections["data"]; !ok {
		return
	}

	if len(data.Pairs) == 0 {
		return
	}

	this.Datastore = make(map[string]string)
	for k, v := range data.Pairs {
		this.Datastore[k] = v
	}
	return
}

func (this *Config) Save(file string) (err os.Error) {
	cfg := ini.NewConfig()
	cfg.AddComment("data",
		`Any values needed to create a valid connection to the db of your choice,
should be added in this section as key/value pairs. For example:

  user = bob
  pass = 1234
  dbname = mudkipz
  dbhost = 127.0.0.1

Refer to the README of the individual db driver for the required keys.`)

	for k, v := range this.Datastore {
		cfg.Set("data", k, v)
	}

	return ini.Save(file, cfg)
}
