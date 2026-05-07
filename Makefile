BINARY=reminders-import

build:
	go build -o bin/$(BINARY) ./cmd

install:
	cp bin/$(BINARY) /usr/local/bin/

clean:
	rm -rf bin/