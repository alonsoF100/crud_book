include .env
export

GOOSE = goose -dir internal/storage/postgres/migrations postgres "${LOCAL_DATABASE_URL}"

# DOCKER COMMANDS

# Start all services in Docker (build + run)
docker-up:
	cd deployments/dev && docker-compose up --build

# Stop and remove all containers (database data preserved)
docker-down:
	cd deployments/dev && docker-compose down

# View logs of all services in real time
docker-logs:
	cd deployments/dev && docker-compose logs -f

# Build images only without starting
docker-build:
	cd deployments/dev && docker-compose build

# DATABASE MIGRATIONS

# Show status of all migrations (which are applied, which are not)
migration-status:
	$(GOOSE) status

# Apply all new migrations
migration-up:
	$(GOOSE) up

# Rollback the last migration
migration-down:
	$(GOOSE) down

# Create new migration (usage: make migration-create name=add_column)
migration-create:
	$(GOOSE) create $(name) sql

# LOCAL DEVELOPMENT

# Run application locally (for fast development)
run:
	go run cmd/api/main.go

# Run all tests in project
test:
	go test ./...

# Build application binary
build:
	go build -o app cmd/api/main.go

# CLEANUP

# Full cleanup: remove binary and all Docker containers/volumes
clean:
	rm -f app
	docker-compose -f deployments/dev/docker-compose.yml down -v

# DATABASE BACKUP

# Create database backup (requires running PostgreSQL)
backup:
	pg_dump -U postgres -d library > backup_$(shell date +%Y%m%d_%H%M%S).sql

# Restore database from backup
restore:
	psql -U postgres -d library < $(file)