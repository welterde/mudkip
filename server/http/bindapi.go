package main

import "os"

func BindApi(methods *ServiceMethodList) (err os.Error) {
	methods.Add(NewServiceMethod("", getHandler))

	sm := NewServiceMethod("", getHandler)
	sm.Method = HEAD
	methods.Add(sm)

	sm = NewServiceMethod("", postHandler)
	sm.Method = POST
	methods.Add(sm)

	sm = NewServiceMethod("", notImplementedHandler)
	sm.Method = CONNECT
	methods.Add(sm)

	sm = NewServiceMethod("", notImplementedHandler)
	sm.Method = DELETE
	methods.Add(sm)

	sm = NewServiceMethod("", notImplementedHandler)
	sm.Method = OPTIONS
	methods.Add(sm)

	sm = NewServiceMethod("", notImplementedHandler)
	sm.Method = PUT
	methods.Add(sm)

	sm = NewServiceMethod("", notImplementedHandler)
	sm.Method = TRACE
	methods.Add(sm)

	return methods.Build()
}
