MODULE=github.com/mehdieidi/carify/data

MIGRATIONS_PATH=internal/database/migrations

DB_NAME=carify_db
DB_URL=postgresql://postgres:1234@localhost:5440/$(DB_NAME)?sslmode=disable

POSTGRES_CONTAINER=carify_postgres

tidy:
	go mod tidy

fmt:
	go fmt ./...
	swag fmt

doc:
	swag i --pd 

build:
	go build -tags postgres

fetch_build:
	./data -fetch

preprocess_build:
	./data -preprocess

onehot_build:
	./data -onehot

postgres_container:
	docker run --name $(POSTGRES_CONTAINER) -p 5440:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1234 -d postgres:alpine

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
	migrate -database "postgresql://postgres:1234@localhost:5440/carify_db?sslmode=disable" -path "internal/database/migrations" force $(VERSION)

fetch: tidy doc fmt build fetch_build

preprocess: tidy doc fmt build preprocess_build

onehot: tidy doc fmt build onehot_build

.PHONY: tidy fmt doc build server postgres_container create_postgres_db drop_postgres_db migrateup migrateup_1 migratedown migratedown_1 create_migration fix_migrate pipeline
