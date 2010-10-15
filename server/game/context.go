package main

import "os"
import "os/signal"
import "sync"
import "log"
import "fmt"
import "mudkip/lib"
import "mudkip/store"

type Context struct {
	config *Config
	world  *lib.WorldInfo
	users  map[string]*lib.User
	lock   *sync.Mutex
	server *Server
	log    *log.Logger
}

func NewContext(cfg *Config, info *lib.WorldInfo) *Context {
	c := new(Context)
	c.config = cfg
	c.lock = new(sync.Mutex)
	c.users = make(map[string]*lib.User)
	c.server = NewServer(cfg.MaxClients, cfg.ClientTimeout, func(m lib.Message) { c.handleMessage(m) })

	var logtarget *os.File

	if cfg.LogFile != "" {
		var err os.Error
		if logtarget, err = os.Open(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0); err != nil {
			logtarget = os.Stdout
		}
	} else {
		logtarget = os.Stdout
	}

	c.log = log.New(logtarget, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	return c
}

func (this *Context) handleMessage(msg lib.Message) {
	var err os.Error

	id := msg.Sender().String()
	this.Info("%s -> %T", id, msg)

	switch tt := msg.(type) {
	case *lib.ClientConnected:
		ds := store.New()
		ds.Open(this.config.Datastore)

		this.lock.Lock()
		this.users[id] = lib.NewUser(ds)
		this.users[id].Info.Zone = this.world.DefaultZone
		this.lock.Unlock()

	case *lib.ClientDisconnected:
		this.lock.Lock()
		if _, ok := this.users[id]; ok {
			this.users[id].Dispose()
			this.users[id] = nil, false
		}
		this.lock.Unlock()

	case *lib.Login:
		client := this.server.GetClient(id)
		if err = this.users[id].Login(tt.Username, tt.Password); err != nil {
			client.Err(err)
		} else {
			client.Ack()
		}

	case *lib.Logout:
		client := this.server.GetClient(id)
		if err = this.users[id].Logout(); err != nil {
			client.Err(err)
		} else {
			client.Ack()
		}

	case *lib.Register:
		client := this.server.GetClient(id)
		if err = this.users[id].Register(tt.Username, tt.Password); err != nil {
			client.Err(err)
		} else {
			client.Ack()
		}
	}
}

func (this *Context) Dispose() {
	this.Info("Shutting down...")

	this.lock.Lock()
	defer this.lock.Unlock()

	if this.server != nil {
		this.server.Close()
		this.server = nil
	}

	for _, usr := range this.users {
		usr.Dispose()
	}
}

func (this *Context) Run() (err os.Error) {
	if err = this.server.Open(
		this.config.ListenAddr,
		this.config.Secure,
		this.config.ServerCert,
		this.config.ServerKey,
	); err != nil {
		return
	}

	this.Info("Listening on: %s", this.config.ListenAddr)
	this.Info("Max clients: %d", this.config.MaxClients)
	this.Info("Client timeout: %d minute(s)", this.config.ClientTimeout)
	this.Info("Secure connection: %v", this.config.Secure)

	var sig signal.Signal

loop:
	for {
		select {
		case sig = <-signal.Incoming:
			if unix, ok := sig.(signal.UnixSignal); ok {
				switch unix {
				case signal.SIGINT, signal.SIGTERM, signal.SIGKILL:
					return
				}
			}
		}

		if closed(signal.Incoming) {
			return
		}
	}

	return
}

func (this *Context) Info(f string, a ...interface{}) {
	s := fmt.Sprintf(f, a...) // issue 1136
	this.lock.Lock()
	this.log.Logf("[i] %s", s)
	this.lock.Unlock()
}

func (this *Context) Error(f string, a ...interface{}) {
	s := fmt.Sprintf(f, a...) // issue 1136
	this.lock.Lock()
	this.log.Logf("[e] %s", s)
	this.lock.Unlock()
}

func (this *Context) Warn(f string, a ...interface{}) {
	s := fmt.Sprintf(f, a...) // issue 1136
	this.lock.Lock()
	this.log.Logf("[w] %s", s)
	this.lock.Unlock()
}

func (this *Context) Debug(f string, a ...interface{}) {
	s := fmt.Sprintf(f, a...) // issue 1136
	this.lock.Lock()
	this.log.Logf("[d] %s", s)
	this.lock.Unlock()
}
