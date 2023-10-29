# Мэйкфайл для управления проектом MyCameraApp

# Переменные
APP_NAME := mycameraapp
DB_DSN := postgres://mycameraapp:mycamera@localhost/mycameraapp?sslmode=disable

# Команды
migrate-up:
	migrate -path=./migrations -database="$(DB_DSN)" up

migrate-down:
	migrate -path=./migrations -database="$(DB_DSN)" down

run:
	go run main.go

# Задачи
.PHONY: migrate-up migrate-down run
