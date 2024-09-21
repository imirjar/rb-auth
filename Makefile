include .env

.PHONY: run build test

APP_FILE=cmd/main.go

docs:
	swag init -g internal/gateway/http/http.go

test: 
	go test ./...
	
run:
	go run $(APP_FILE)

build:
	go build -o bin/rb_auth $(APP_FILE)

start: docs test build
	./bin/rb_auth

