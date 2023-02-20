build:
	go build -o bin/main cmd/main.go

run: build
	./bin/main

dev:
	nodemon --exec go run cmd/main.go --signal SIGTERM