include .env

.PHONY: run build

APP_FILE=cmd/main.go

run:
	go run $(APP_FILE)

build:
	go build -o bin/rb_auth $(APP_FILE)

build_and_run: build
	./bin/rb_auth