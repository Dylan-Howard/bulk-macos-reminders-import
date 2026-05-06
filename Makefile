BINARY=reminders-import

build:
	go build -o bin/$(BINARY) ./cmd/reminders-import

install:
	cp bin/$(BINARY) /usr/local/bin/

clean:
	rm -rf bin/