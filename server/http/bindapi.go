package main

import "os"

func BindApi(methods *ServiceMethodList) (err os.Error) {
	// TODO: Bind Mudkip API

	// Catch-all handlers for HTTP commands we have not yet covered.
	methods.Add(NewServiceMethod("", getHandler, GET))
	methods.Add(NewServiceMethod("", getHandler, HEAD))
	methods.Add(NewServiceMethod("", postHandler, POST))
	methods.Add(NewServiceMethod("", notImplementedHandler, CONNECT))
	methods.Add(NewServiceMethod("", notImplementedHandler, DELETE))
	methods.Add(NewServiceMethod("", notImplementedHandler, OPTIONS))
	methods.Add(NewServiceMethod("", notImplementedHandler, PUT))
	methods.Add(NewServiceMethod("", notImplementedHandler, TRACE))
	return methods.Build()
}
