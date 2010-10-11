
all:
	make -C lib install
	make -C data
	make -C server
	make -C client

clean:
	make -C lib clean
	make -C data clean
	make -C server clean
	make -C client clean

test:
	make -C lib test
	make -C data test

format:
	gofmt -w .
