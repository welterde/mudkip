package main

import "os"
import "ini"

type Config struct {
	Server            string
	Secure            bool
	AcceptInvalidCert bool
}

func NewConfig() *Config {
	c := new(Config)
	c.Secure = false
	c.AcceptInvalidCert = false
	return c
}

func (this *Config) Load(file string) (err os.Error) {
	var cfg *ini.Config
	if cfg, err = ini.Load(file); err != nil {
		return
	}

	this.Server = cfg.S("net", "address", "")
	this.Secure = cfg.B("net", "secure", false)
	this.AcceptInvalidCert = cfg.B("net", "acceptinvalidcert", false)
	return
}

func (this *Config) Save(file string) (err os.Error) {
	cfg := ini.NewConfig()
	cfg.AddComment("net", "Address should be in the format ip:port. It can be in")
	cfg.AddComment("net", "IPv4 and IPv6 format. IPv6 address should be encased")
	cfg.AddComment("net", "in brackets. For example:")
	cfg.AddComment("net", "  address = 127.0.0.1:54321")
	cfg.AddComment("net", "  address = [::1]:54321")
	cfg.AddComment("net", "  address = :54321")
	cfg.AddComment("net", "")
	cfg.AddComment("net", "acceptinvalidcert determines if we are willing to accept self-signed")
	cfg.AddComment("net", "certificates or not. These are not verified by an official CA, but setting")
	cfg.AddComment("net", "this to true is good enough for debugging, or when you trust the source")
	cfg.AddComment("net", "implicitely.")

	cfg.Set("net", "address", this.Server)
	cfg.Set("net", "secure", this.Secure)
	cfg.Set("net", "acceptinvalidcert", this.AcceptInvalidCert)
	return ini.Save(file, cfg)
}
