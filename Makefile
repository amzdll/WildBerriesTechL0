include .env

GOOSE_QUERY=${LOCAL_BIN}/goose -dir ${PG_MIGRATION_DIR} postgres ${PG_DSN}

install-deps:
	GOBIN=${LOCAL_BIN} go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=${LOCAL_BIN} go install github.com/swaggo/swag/cmd/swag@latest

local-migration-status:
	$(GOOSE_QUERY) status -v

local-migration-up:
	$(GOOSE_QUERY) up -v

local-migration-down:
	$(GOOSE_QUERY) down -v

generate-swagger:
	${LOCAL_BIN}/swag init -g ${SWAGGER_SRC_DIR} -o ${SWAGGER_OUTPUT_DIR}




