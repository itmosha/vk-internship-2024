
include .env

build:
	go build -v -o filmsbinary.out ./cmd/app/main.go

build_stripped:
	go build -ldflags="-w -s" -o filmsbinary.out ./cmd/app/main.go

makemigration:
	migrate create -ext sql -dir migrations $(name)

migrate_up:
	migrate -path migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable" -verbose up

migrate_down:
	migrate -path migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable" -verbose down

.DEFAULT_GOAL := build