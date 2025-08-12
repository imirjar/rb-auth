include .env

.PHONY: run build test prod docker-build docker-run docker-clean


APP_FILE=cmd/main.go

docs:
	swag init -g internal/gateway/http/http.go 

test: 
	go test ./...

run:
	go run $(APP_FILE)

build:
	go build -o rb_auth $(APP_FILE)

start: docs test build 
	./bin/rb_auth

# Docker сборка
docker-build: 
	docker build -t rb_auth .


# Запуск проекта в Docker контейнере с использованием .env файла
docker-run: docker-build
	docker run -d --name rb_auth -p 7070:7070 rb_auth

# Очищение контейнера после остановки
docker-clean:
	@if [ "$(shell docker ps -a -q --filter "name=rb_auth")" ]; then \
		docker rm rb_auth; \
	fi

# Полный цикл сборки и запуска в Docker
prod: docker-clean docker-run