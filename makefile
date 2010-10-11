
all:
	make -C lib install
	make -C data/mongo install
	make -C data/mysql install
	make -C data/sqlite install
	make -C server
	make -C client

clean:
	make -C lib clean
	make -C data/mongo clean
	make -C data/mysql clean
	make -C data/sqlite clean
	make -C server clean
	make -C client clean

test:
	make -C lib test
	make -C data/mongo test
	make -C data/mysql test
	make -C data/sqlite test

format:
	gofmt -w .
