DB_NAME ?= cushon
DB_USER ?= cushon
DB_PASS ?= mypassword
DB_HOST ?= localhost
DB_PORT ?= 5432

envars = DB_NAME=$(DB_USER) \
	DB_USER=$(DB_USER) \
	DB_PASS=$(DB_PASS) \
	DB_HOST=$(DB_HOST) \
	DB_PORT=$(DB_PORT) \

# https://go.dev/blog/vuln
.PHONY: audit
audit:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

# https://github.com/golangci/golangci-lint/blob/master/.golangci.yml
.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: logs-api
logs-api: 
	docker-compose logs api

.PHONY: run-api-local
run-api-local:
	cd cmd/assignment; $(envars) go run handler.go router.go main.go

.PHONY: run-db
run-db: 
	docker-compose up -d postgres

.PHONY: start
start:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down -v

.PHONY: test/unit
test/unit:
	go test -tags unit ./...

check: audit lint test/unit
start-local: run-db run-api-local
