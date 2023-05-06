ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=gojhw5 host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MIGRATION_FOLDER=$(INTERNAL_PKG_PATH)/db/migrations

SCRIPTS_PKG_PATH=$(CURDIR)/scripts
INSERT_DATA_PATH=$(SCRIPTS_PKG_PATH)/insert_data.sql

.PHONY: insert-mock-data
insert-mock-data:
	psql -h localhost -d gojhw5 -U test -p 5432 -a -q -f "$(INSERT_DATA_PATH)"

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: test-clean
test-clean:
	go clean -testcache

.PHONY: test
test:
	go test ./...

.PHONY: test-integration
test-integration:
	go test ./... -tags=integration

.PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: compose-up
compose-up:
	docker-compose build
	docker-compose up

.PHONY: compose-rm
compose-rm:
	docker-compose down

.PHONY: generate
generate:
	rm -rf internal/app/pb
	mkdir -p internal/app/pb

	protoc \
		--proto_path=proto/ \
		--go_out=internal/app/pb \
		--go-grpc_out=internal/app/pb \
		proto/*.proto