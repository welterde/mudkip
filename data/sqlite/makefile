include $(GOROOT)/src/Make.inc

TARG = mudkip/store
GOFILES = store.go
CGOFILES = sqlite.go

ifeq ($(GOOS),darwin)
CGO_LDFLAGS=/usr/lib/libsqlite3.0.dylib
else
CGO_LDFLAGS=-lsqlite3
endif

include $(GOROOT)/src/Make.pkg
