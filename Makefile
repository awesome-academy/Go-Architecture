include .env
# Directory containing the Makefile.
PROJECT_ROOT = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Directories containing independent Go modules.
MODULE_DIRS = .

.PHONY: all
all: lint

.PHONY: lint
lint: golangci-lint

.PHONY: golangci-lint
golangci-lint:
	@$(foreach mod,$(MODULE_DIRS), \
		(cd $(mod) && \
		echo "[lint] golangci-lint: $(mod)" && \
		golangci-lint run --path-prefix $(mod)) &&) true

.PHONY: migrateup
migrateup: 
	migrate -path db/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

.PHONY: migratedown
migratedown: 
	migrate -path db/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down