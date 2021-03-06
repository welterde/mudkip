package main

import "os"
import "ini"

type Config struct {
	ListenAddr    string
	Secure        bool
	ServerCert    string
	ServerKey     string
	CookieSalt    string
	WebRoot       string
	ServerName    string
	ClientTimeout int
	Datastore     map[string]string
}

func NewConfig() *Config {
	c := new(Config)
	c.Secure = false
	c.ClientTimeout = 2
	c.ServerCert = "/path/to/cert.pem"
	c.ServerKey = "/path/to/key.pem"
	c.Datastore = make(map[string]string)
	c.CookieSalt = "xxxx"
	c.ServerName = "MUDkip"
	c.WebRoot = "webroot/"
	return c
}

func (this *Config) Load(file string) (err os.Error) {
	var cfg *ini.Config
	if cfg, err = ini.Load(file); err != nil {
		return
	}

	this.ListenAddr = cfg.S("net", "address", "")
	this.Secure = cfg.B("net", "secure", false)
	this.ServerCert = cfg.S("net", "servercert", "/path/to/cert.pem")
	this.ServerKey = cfg.S("net", "serverkey", "/path/to/key.pem")
	this.ClientTimeout = cfg.I("net", "clienttimeout", 2)
	this.CookieSalt = cfg.S("net", "cookiesalt", "xxxx")
	this.WebRoot = cfg.S("net", "webroot", "webroot/")
	this.ServerName = cfg.S("net", "servername", "MUDkip")

	var data *ini.Section
	var ok bool

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
	cfg.AddComment("net",
		`- address should be in the format ip:port. It can be in IPv4 and IPv6 format.
  IPv6 address should be encased in brackets. The ip can be omitted. This will
  automatically bind the server to localhost on the given port. For example:
 
    address = 127.0.0.1:54321
    address = [::1]:54321
    address = :54321
 
- servercert and serverkey must be set when secure = true.
- servername is added to a HTTP response in the Server header.
- cookiesalt is a string of random character used to encrypt cookies.
  Keep this value from becoming public.`)

	cfg.Set("net", "address", this.ListenAddr)
	cfg.Set("net", "secure", this.Secure)
	cfg.Set("net", "servercert", this.ServerCert)
	cfg.Set("net", "serverkey", this.ServerKey)
	cfg.Set("net", "clienttimeout", this.ClientTimeout)
	cfg.Set("net", "cookiesalt", this.CookieSalt)
	cfg.Set("net", "webroot", this.WebRoot)
	cfg.Set("net", "servername", this.ServerName)

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
