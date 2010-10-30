package main

import "sync"
import "time"
import "fmt"
import "mudkip/lib"

type Context struct {
	config   *Config
	lock     *sync.RWMutex
	sessions map[string]*Session
	worlds   []*lib.World
}

func NewContext(config *Config) *Context {
	c := new(Context)
	c.config = config
	c.lock = new(sync.RWMutex)
	c.sessions = make(map[string]*Session)
	c.worlds = make([]*lib.World, 0, 32)
	return c
}

func (this *Context) Worlds() []*lib.World { return this.worlds }
func (this *Context) Config() *Config      { return this.config }
func (this *Context) Sessions() []*Session {
	this.lock.Lock()
	defer this.lock.Unlock()

	var idx int
	list := make([]*Session, len(this.sessions))

	for _, v := range this.sessions {
		list[idx] = v
		idx++
	}

	return list
}

func (this *Context) CreateSession(addr string) *Session {
	id := fmt.Sprintf("%s%d", addr, time.Nanoseconds())

	this.lock.Lock()
	defer this.lock.Unlock()

	this.sessions[id] = NewSession(id)
	return this.sessions[id]
}

func (this *Context) GetSession(id string) *Session {
	this.lock.Lock()
	defer this.lock.Unlock()

	if session, ok := this.sessions[id]; ok {
		return session
	}

	return nil
}

func (this *Context) AddWorld(v *lib.World) {
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
