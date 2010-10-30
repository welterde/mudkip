package main

import "os"
import "strings"
import "sync"

type ServiceMethodList struct {
	l []*ServiceMethod
	m *sync.RWMutex
}

func NewServiceMethodList() *ServiceMethodList {
	rc := new(ServiceMethodList)
	rc.l = make([]*ServiceMethod, 0)
	rc.m = new(sync.RWMutex)
	return rc
}

func (this *ServiceMethodList) Build() (err os.Error) {
	this.m.Lock()
	defer this.m.Unlock()

	for _, sm := range this.l {
		if err = sm.Build(); err != nil {
			return
		}
	}

	return
}

func (this *ServiceMethodList) Exec(sc *ServiceContext) bool {
	this.m.Lock()
	defer this.m.Unlock()

	for _, sm := range this.l {
		if !sm.Method.Equals(sc.Req.Method) {
			continue
		}

		if sc.Params = sm.pattern.FindStringSubmatch(sc.Req.URL.Path); len(sc.Params) == 0 {
			continue
		}

		sc.Params = sc.Params[1:]
		return sm.handler(sc)
	}

	return false
}

func (this *ServiceMethodList) Add(sm *ServiceMethod) {
	if idx := this.Indexof(sm.Method, sm.Name); idx != -1 {
		this.l[idx] = sm
		return
	}

	this.m.Lock()
	defer this.m.Unlock()

	sz := len(this.l)

	if sz >= cap(this.l) {
		cp := make([]*ServiceMethod, sz, sz+10)
		copy(cp, this.l)
		this.l = cp
	}

	this.l = this.l[0 : sz+1]
	this.l[sz] = sm
}

func (this *ServiceMethodList) Remove(sm *ServiceMethod) {
	idx := this.Indexof(sm.Method, sm.Name)
	if idx == -1 {
		return
	}

	this.m.Lock()
	defer this.m.Unlock()

	this.l[idx].pattern = nil
	this.l[idx].handler = nil
	this.l[idx] = nil

	copy(this.l[idx:], this.l[idx+1:])
	this.l = this.l[0 : len(this.l)-1]
}

func (this *ServiceMethodList) Clear() {
	this.m.Lock()
	defer this.m.Unlock()

	for i, _ := range this.l {
		this.l[i] = nil
	}

	this.l = nil
}

func (this *ServiceMethodList) Indexof(method HttpMethod, name string) int {
	for i, sm := range this.l {
		if sm.Method == method && sm.Name == name {
			return i
		}
	}

	return -1
}

func (this *ServiceMethodList) Find(method HttpMethod, name string) *ServiceMethod {
	for _, sm := range this.l {
		if sm.Method == method && sm.Name == name {
			return sm
		}
	}

	return nil
}

func (this *ServiceMethodList) FindFiltered(name string) []*ServiceMethod {
	name = strings.ToLower(name)
	all := len(name) == 0
	list := make([]*ServiceMethod, 0, len(this.l))
	sz := 0

	for _, v := range this.l {
		if len(v.Name) > 0 && (all || strings.HasPrefix(strings.ToLower(v.Name), name)) {
			sz = len(list)
			list = list[0 : sz+1]
			list[sz] = v
		}
	}

	return list
}

func (this *ServiceMethodList) Contains(method HttpMethod, name string) bool {
	for _, sm := range this.l {
		if sm.Method == method && sm.Name == name {
			return true
		}
	}

	return false
}
