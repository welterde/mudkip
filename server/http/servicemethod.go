package main

import "os"
import "regexp"
import "bytes"

type ServiceHandler func(*ServiceContext) bool

type ServiceMethod struct {
	Method  HttpMethod
	Name    string
	Params  []string
	pattern *regexp.Regexp
	handler ServiceHandler
}

func NewServiceMethod(name string, sh ServiceHandler) *ServiceMethod {
	sm := new(ServiceMethod)
	sm.Name = name
	sm.handler = sh
	sm.Method = GET
	sm.Params = []string{}
	return sm
}

func (this *ServiceMethod) Build() (err os.Error) {
	var data []byte
	buf := bytes.NewBuffer(data)

	if len(this.Name) == 0 {
		buf.WriteString("^/.*$")
	} else {
		buf.WriteString("^/")
		buf.WriteString(this.Name)
		buf.WriteRune('/')

		for _, p := range this.Params {
			buf.WriteString(p)
			buf.WriteRune('/')
		}
	}

	data = buf.Bytes()
	data[len(data)-1] = '$'

	this.pattern, err = regexp.Compile(string(data))
	return
}
