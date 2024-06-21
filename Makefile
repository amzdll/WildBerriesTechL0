include .env

LOCAL_BIN=$(PWD)/bin
GOOSE_QUERY=${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN}
PG_MIGRATION_DIR=${PWD}/migrations
SWAGGER_SRC_DIR=./internal/api/swagger/handler.go
SWAGGER_OUTPUT_DIR=api/

.PHONY: install-deps
install-deps:
	GOBIN=${LOCAL_BIN} go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=${LOCAL_BIN} go install github.com/swaggo/swag/cmd/swag@latest
	GOBIN=${LOCAL_BIN} go install golang.org/x/tools/cmd/goimports
	GOBIN=${LOCAL_BIN} go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up

.PHONY: send-message
send-message:
	@sh scripts/send_data.sh

.PHONY: local-migration-status
local-migration-status:
	$(GOOSE_QUERY) status -v

.PHONY: local-migration-up
local-migration-up:
	$(GOOSE_QUERY) up -v

.PHONY: local-migration-down
local-migration-down:
	$(GOOSE_QUERY) down -v

.PHONY: generate-swagger
generate-swagger:
	${LOCAL_BIN}/swag init -g ${SWAGGER_SRC_DIR} -o ${SWAGGER_OUTPUT_DIR}

.PHONY: style
style: install-deps
	${LOCAL_BIN}/golangci-lint run



