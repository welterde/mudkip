
all:
	make -C lib install
	make -C server
	make -C client

clean:
	make -C lib clean
	make -C server clean
	make -C client clean

test:
	make -C lib test

format:
	gofmt -w .
