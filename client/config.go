package main

import "os"
import "ini"

type Config struct {
	Server string
	Secure bool
}

func NewConfig() *Config {
	c := new(Config)
	c.Secure = false
	return c
}

func (this *Config) Load(file string) (err os.Error) {
	var cfg *ini.Config
	if cfg, err = ini.Load(file); err != nil {
		return
	}

	this.Server = cfg.S("net", "address", "")
	this.Secure = cfg.B("net", "secure", false)
	return
}

func (this *Config) Save(file string) (err os.Error) {
	cfg := ini.NewConfig()
	cfg.Set("net", "address", this.Server)
	cfg.Set("net", "secure", this.Secure)
	return ini.Save(file, cfg)
}
