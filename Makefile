include .env
export

.PHONY: run fmt test tidy \
	migrate-create migrate-up migrate-down migrate-status migrate-reset

run:
	go run ./cmd/hermes

fmt:
	go fmt ./...

test:
	go test ./...

tidy:
	go mod tidy

migrate-create:
ifndef name
	$(error Usage: make migrate-create name=create_users)
endif
	goose -dir $(GOOSE_MIGRATION_DIR) create $(name) sql

migrate-up:
	goose up

migrate-down:
	goose down

migrate-status:
	goose status

migrate-reset:
	goose reset