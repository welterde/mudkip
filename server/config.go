package main

import "os"
import "ini"

type Config struct {
	ListenAddr string
	Secure     bool
	MaxClients int
	LogFile    string
}

func NewConfig() *Config {
	c := new(Config)
	c.MaxClients = 16
	c.Secure = false
	return c
}

func (this *Config) Load(file string) (err os.Error) {
	var cfg *ini.Config
	if cfg, err = ini.Load(file); err != nil {
		return
	}

	this.ListenAddr = cfg.S("net", "address", "")
	this.Secure = cfg.B("net", "secure", false)
	this.MaxClients = cfg.I("net", "maxclients", 16)
	this.LogFile = cfg.S("misc", "logfile", "")
	return
}

func (this *Config) Save(file string) (err os.Error) {
	cfg := ini.NewConfig()
	cfg.Set("net", "address", this.ListenAddr)
	cfg.Set("net", "secure", this.Secure)
	cfg.Set("net", "maxclients", this.MaxClients)
	cfg.Set("misc", "logfile", this.LogFile)
	return ini.Save(file, cfg)
}
