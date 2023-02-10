all: build

build:
	go build -o server ./server.go
	go build -o client ./client.go

run-server:
	./server

run-client:
	./client

clean:
	rm server client

