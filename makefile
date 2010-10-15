
# Set this to whatever datastore support you want to compile into the server.
# For a listing of all available drivers, see mudkip/data/
DATASTORE = sqlite

all:
	make -C builder install
	make -C lib install
	make -C data/$(DATASTORE) install
	make -C dsinit
	make -C server
	make -C client

clean:
	make -C builder clean
	make -C lib clean
	make -C data/$(DATASTORE) clean
	make -C dsinit clean
	make -C server clean
	make -C client clean

test:
	make -C builder test
	make -C lib test
	make -C data/$(DATASTORE) test

format:
	gofmt -w .
