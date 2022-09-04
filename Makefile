.PHONY: up down pull update clean install migrate test generate-sqlboiler generate-mocks psql run

up:
	docker compose up --detach postgres

down:
	docker compose down --remove-orphans

pull:
	docker compose pull

update: pull up
	docker system prune

clean:
	docker compose down --remove-orphans --volumes

install:
	go install github.com/rubenv/sql-migrate/...@v1.1.2

migrate:
	sql-migrate up

test:
	go test -cover ./...

generate-sqlboiler:
	sqlboiler psql
	goimports -w models/*.go

generate-mocks:
	docker compose run --rm mockery --recursive=false --with-expecter=true --dir=postgres --name=Postgres

psql:
	docker compose exec postgres psql -U test

run:
	go run ./cmd/api
