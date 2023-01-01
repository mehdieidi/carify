MODULE=back

MIGRATIONS_PATH=internal/database/migrations

DB_NAME=back_test
DB_URL=postgresql://postgres:1234@localhost:5432/$(DB_NAME)?sslmode=disable

POSTGRES_CONTAINER=postgres_carify
REDIS_CONTAINER=redis_carify

tidy:
	go mod tidy

fmt:
	go fmt ./...
	swag fmt

lint:
	golangci-lint run --config "./config/.golangci.yml"

sort_imports:
	goimportssort -v -w -local $(MODULE) . 

doc:
	swag i --pd 

build:
	go build -tags postgres

server:
	./back

postgres_container:
	docker run --name $(POSTGRES_CONTAINER) -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1234 -d postgres:alpine

redis_container:
	docker run --name $(REDIS_CONTAINER) -p 6379:6379 -e REDIS_PASSWORD=redis -d redis:alpine

create_postgres_db:
	docker exec -it $(POSTGRES_CONTAINER) createdb --username=postgres --owner=postgres $(DB_NAME)

drop_postgres_db:
	docker exec -it $(POSTGRES_CONTAINER) dropdb --username=postgres $(DB_NAME)

migrateup:
	migrate -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" -verbose up

migrateup_1:
	migrate -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" -verbose down

migratedown_1:
	migrate -path "$(MIGRATIONS_PATH)" -database "$(DB_URL)" -verbose down 1

create_migration:
	migrate create -ext sql -dir "$(MIGRATIONS_PATH)" -seq "$(MIGRATION_NAME)"

fix_migrate:
	migrate -database "postgresql://postgres:1234@localhost:5432/tavanito_test?sslmode=disable" -path "internal/database/migrations" force $(VERSION)

pipeline: tidy doc fmt build server

.PHONY: tidy fmt lint sort_imports doc build server postgres_container redis_container create_postgres_db drop_postgres_db migrateup migrateup_1 migratedown migratedown_1 create_migration fix_migrate pipeline
