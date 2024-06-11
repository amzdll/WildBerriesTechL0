include .env

GOOSE_QUERY=${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN}

install-deps:
	# Migration
	GOBIN=${LOCAL_BIN} go install github.com/pressly/goose/v3/cmd/goose@latest

migration-status:
	$(GOOSE_QUERY) status -v

migration-up:
	$(GOOSE_QUERY) up -v

migration-down:
	$(GOOSE_QUERY) down -v


