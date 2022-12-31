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
	sql-migrate up

seed: migrate
	docker compose cp ./db/seed.sql postgres:/tmp
	docker compose exec postgres psql -U test -f /tmp/seed.sql

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
