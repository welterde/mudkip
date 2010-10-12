package main

import "os"
import "os/signal"
import "sync"
import "log"
import "fmt"
import "mudkip/lib"

type Context struct {
	config *Config
	users  map[string]*User
	lock   *sync.Mutex
	server *Server
	log    *log.Logger
}

func NewContext(cfg *Config) *Context {
	c := new(Context)
	c.config = cfg
	c.lock = new(sync.Mutex)
	c.users = make(map[string]*User)
	c.server = NewServer(cfg.MaxClients, cfg.ClientTimeout)

	var logtarget *os.File
	if cfg.LogFile != "" {
		var err os.Error
		if logtarget, err = os.Open(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0); err != nil {
			logtarget = os.Stdout
		}
	} else {
		logtarget = os.Stdout
	}

	c.log = log.New(logtarget, nil, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	return c
}

func (this *Context) Dispose() {
	this.Info("Shutting down...")

	this.lock.Lock()
	defer this.lock.Unlock()

	if this.server != nil {
		this.server.Close()
		this.server = nil
	}
}

func (this *Context) HandleMessage(msg lib.Message) {
	switch tt := msg.(type) {
	case *lib.ClientConnected:
		this.Info("Client connected: %s", msg.Sender())
		this.lock.Lock()

		this.lock.Unlock()

	case *lib.ClientDisconnected:
		this.Info("Client disconnected: %s", msg.Sender())
		this.lock.Lock()

		this.lock.Unlock()
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

	this.Info("Listening on: %s", this.server.conn.Addr())
	this.Info("Max clients: %d", this.config.MaxClients)
	this.Info("Client timeout: %d minute(s)", this.config.ClientTimeout)
	this.Info("Secure connection: %v", this.config.Secure)

	var msg lib.Message
	var sig signal.Signal

	incoming := this.server.Messages()

loop:
	for {
		select {
		case msg = <-incoming:
			go this.HandleMessage(msg)

		case sig = <-signal.Incoming:
			if unix, ok := sig.(signal.UnixSignal); ok {
				switch unix {
				case signal.SIGINT, signal.SIGTERM, signal.SIGKILL:
					return
				}
			}
		}

		if closed(incoming) || closed(signal.Incoming) {
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
