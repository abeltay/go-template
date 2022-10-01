.PHONY: up down pull update clean migrate test generate-sqlboiler generate-mocks psql run

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

migrate:
	go run ./cmd/migrate

test:
	go test -cover ./...

install-sqlboiler:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/volatiletech/sqlboiler/v4@v4.12.0
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.12.0

generate-sqlboiler:
	sqlboiler psql
	goimports -w models/*.go

generate-mocks:
	docker compose run --rm mockery --recursive=false --with-expecter=true --dir=postgres --name=Postgres

psql:
	docker compose exec postgres psql -U test

run:
	go run ./cmd/api
