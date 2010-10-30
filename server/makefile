include $(GOROOT)/src/Make.inc

TARG = server
GOFILES = main.go util.go config.go context.go http/*.go

all:
	$(GC) -o $(TARG).6 $(GOFILES)
	$(LD) -s -o $(TARG) $(TARG).6

clean:
	rm -rf *.o *.a *.[568vq] [568vq].out *.cgo1.go *.cgo2.c _cgo_defun.c _cgo_gotypes.go _cgo_export.* *.so *.exe $(TARG)
