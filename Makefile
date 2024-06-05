include scripts/common.mk

dc-up: ## Docker compose up
	@docker compose up -d

dc-down: ## Docker compose down
	@docker compose down

mig-new: ## Create migration script, required name=<example>
ifdef name
	@goose -v create $(name) sql
endif

mig-up: ## Migrate up
	@goose -v up

mig-down: ## Migrate down
	@goose -v down

mig-ver: ## Show current migration version
	@goose -v version
	@goose -v status

gen-query: ## Generate queries
	@sqlc generate

gen-mock: ## Generate mocks
	@mockgen -package mockdb -destination db/mock/store.go github.com/hieulw/gobank/db/sqlc Store

psql: ## Access to postgres database
	@docker compose exec -it postgres psql -U ${DB_USERNAME} -d ${DB_DATABASE}

clean: ## Clean cache, coverage, build outputs
	@rm -f **/**.out # clean output files
	@rm -f main
	@go clean -cache -testcache
	@go mod tidy

lint: ## Run lint
	@golangci-lint run -v

build: ## Build application
	@go build -o main cmd/api/main.go

run: ## Run the application
	@go run cmd/api/main.go

test: ## Test the application
	@go test ./... -v -cover -race

watch: ## Live Reload
	@$(CURDIR)/scripts/air.sh

hurl: ## Test API using hurl
	@hurl --test ./hurls/*

.PHONY: all build run test clean
