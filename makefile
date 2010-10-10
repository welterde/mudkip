
all:
	make -C lib install
	make -C server install
	make -C client install

clean:
	make -C lib clean
	make -C server clean
	make -C client clean

test:
	make -C lib test
	make -C server test
	make -C client test

format:
	gofmt -w .
