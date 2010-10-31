package main

import "sync"
import "time"
import "fmt"
import "mudkip/lib"

type ServerContext struct {
	config   *Config
	lock     *sync.RWMutex
	sessions map[string]*Session
	worlds   []*lib.World
}

func NewServerContext(config *Config) *ServerContext {
	c := new(ServerContext)
	c.config = config
	c.lock = new(sync.RWMutex)
	c.sessions = make(map[string]*Session)
	c.worlds = make([]*lib.World, 0, 32)
	return c
}

func (this *ServerContext) Worlds() []*lib.World { return this.worlds }
func (this *ServerContext) Config() *Config      { return this.config }
func (this *ServerContext) Sessions() []*Session {
	this.lock.RLock()
	defer this.lock.RUnlock()

	var idx int
	list := make([]*Session, len(this.sessions))

	for _, v := range this.sessions {
		list[idx] = v
		idx++
	}

	return list
}

func (this *ServerContext) CreateSession(addr string) *Session {
	this.lock.Lock()
	defer this.lock.Unlock()

	id := fmt.Sprintf("%s%d", addr, time.Nanoseconds())
	this.sessions[id] = NewSession(id)
	return this.sessions[id]
}

func (this *ServerContext) GetSession(id string) *Session {
	this.lock.RLock()
	defer this.lock.RUnlock()

	if session, ok := this.sessions[id]; ok {
		return session
	}

	return nil
}

func (this *ServerContext) AddWorld(v *lib.World) {
	this.lock.Lock()
	defer this.lock.Unlock()

	sz := len(this.worlds)
	if sz >= cap(this.worlds) {
		cp := make([]*lib.World, sz, sz+32)
		copy(cp, this.worlds)
		this.worlds = cp
	}

	this.worlds = this.worlds[0 : sz+1]
	this.worlds[sz] = v
}
