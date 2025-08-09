# Makefile
# Detect OS (Windows or Unix-like)
ifeq ($(OS),Windows_NT)
    RM = del
    DEVNULL = NUL
    SHELL := cmd
else
    RM = rm -f
    DEVNULL = /dev/null
    SHELL := /bin/bash
endif

MIGRATE=migrate
MIGRATION_PATH=./database/migrations
ENV_FILE=.env
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

include $(ENV_FILE)
export

migrate-up:
	@echo Running database migrations...
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" up
	@echo Migration completed.

migrate-down:
	@echo Downing database...
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" down
	@echo Downing completed.

migrate-create:
	@echo Creating migration...
	@$(MIGRATE) create -ext sql -dir $(MIGRATION_PATH) -seq $(name)
	@echo Creating migration completed.

# use for rollback
migrate-goto:
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" goto $(version)

migrate-force:
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" force $(version)
	@echo Successfully force to version.

migrate-drop:
	@$(MIGRATE) -path $(MIGRATION_PATH) -database "$(DB_URL)" drop -f
	@echo Successfully drop all table.

migrate-seed:
	go run cmd/seed/main.go
	@echo Successfully seed

help:
	@echo "Available commands:"
	@echo "  make migrate-up      # Run DB migrations"
	@echo "  make migrate-down    # Rollback last migration"
	@echo "  make migrate-create    # create new migration name="
