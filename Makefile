include .env
export

GOOSE = goose -dir migrations postgres "${DATABASE_URL}"

migration-status:
	$(GOOSE) status

migration-up:
	$(GOOSE) up

migration-down:
	$(GOOSE) down

run:
	go run main.go

test:
	go test ./...