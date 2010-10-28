include $(GOROOT)/src/Make.inc

TARG = mudkip/store
GOFILES = store.go world.go user.go armor.go character.go class.go consumable.go \
	currency.go group.go inventory.go portal.go race.go stats.go weapon.go zone.go
CGOFILES = sqlite.go

ifeq ($(GOOS),darwin)
CGO_LDFLAGS=/usr/lib/libsqlite3.0.dylib
else
CGO_LDFLAGS=-lsqlite3
endif

include $(GOROOT)/src/Make.pkg
