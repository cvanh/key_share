# https://stackoverflow.com/questions/73908737/permission-denied-when-running-go-from-makefile

install: 
	go get;

build: 
	go build -o ./bin/http ./cmd/http;
	go build -o ./bin/ssh ./cmd/ssh;
run: 
	go run ./cmd/http;
	go run ./cmd/ssh;