package main

import "os"
import "ini"
import "strings"
import "mudkip/lib"

type Config struct {
	ListenAddr    string
	Secure        bool
	ServerCert    string
	ServerKey     string
	MaxClients    int
	LogFile       string
	ClientTimeout int
	Datastore     struct {
		Driver string
		Params map[string]string
	}
}

func NewConfig() *Config {
	c := new(Config)
	c.MaxClients = 32
	c.Secure = false
	c.ClientTimeout = 2
	c.ServerCert = "/path/to/cert.pem"
	c.ServerKey = "/path/to/key.pem"
	return c
}

func (this *Config) Load(file string) (err os.Error) {
	var cfg *ini.Config
	if cfg, err = ini.Load(file); err != nil {
		return
	}

	this.LogFile = cfg.S("misc", "logfile", "")
	this.ListenAddr = cfg.S("net", "address", "")
	this.Secure = cfg.B("net", "secure", false)
	this.ServerCert = cfg.S("net", "servercert", "/path/to/cert.pem")
	this.ServerKey = cfg.S("net", "serverkey", "/path/to/key.pem")
	this.MaxClients = cfg.I("net", "maxclients", 32)
	this.ClientTimeout = cfg.I("net", "clienttimeout", 2)

	var data *ini.Section
	var ok bool

	if data, ok = cfg.Sections["data"]; !ok {
		return
	}

	this.Datastore.Driver = data.S("driver", "")
	if len(data.Pairs) <= 1 {
		return
	}

	this.Datastore.Params = make(map[string]string)
	for k, v := range data.Pairs {
		if k != "driver" {
			this.Datastore.Params[k] = v
		}
	}
	return
}

func (this *Config) Save(file string) (err os.Error) {
	cfg := ini.NewConfig()
	cfg.AddComment("net", "Address should be in the format ip:port. It can be in IPv4 and IPv6 format.")
	cfg.AddComment("net", "IPv6 address should be encased in brackets. For example:")
	cfg.AddComment("net", "  address = 127.0.0.1:54321")
	cfg.AddComment("net", "  address = [::1]:54321")
	cfg.AddComment("net", "  address = :54321")
	cfg.AddComment("net", "")
	cfg.AddComment("net", "servercert and serverkey must be set when secure = true")

	cfg.Set("net", "address", this.ListenAddr)
	cfg.Set("net", "secure", this.Secure)
	cfg.Set("net", "servercert", this.ServerCert)
	cfg.Set("net", "serverkey", this.ServerKey)
	cfg.Set("net", "maxclients", this.MaxClients)
	cfg.Set("net", "clienttimeout", this.ClientTimeout)

	cfg.AddComment("data", "The value for 'driver' should be any of the supported driver names: ")
	cfg.AddComment("data", strings.Join(lib.ListStores(), ", "))
	cfg.AddComment("data", "")
	cfg.AddComment("data", "Any additional values needed to create a valid connection to the db of")
	cfg.AddComment("data", "your choice, should be appended in this section as key/value pairs.")
	cfg.AddComment("data", "For example:")
	cfg.AddComment("data", "  driver = mysql")
	cfg.AddComment("data", "  user = bob")
	cfg.AddComment("data", "  pass = 1234")
	cfg.AddComment("data", "  dbname = mudkipz")
	cfg.AddComment("data", "  dbhost = 127.0.0.1")

	cfg.Set("data", "driver", this.Datastore.Driver)

	for k, v := range this.Datastore.Params {
		cfg.Set("data", k, v)
	}

	cfg.AddComment("misc", "The logfile can be left empty if you want the server log to be written to stdout.")
	cfg.Set("misc", "logfile", this.LogFile)
	return ini.Save(file, cfg)
}
