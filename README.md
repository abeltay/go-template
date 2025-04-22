# Go Template

## Setup

1. Install [Go](https://go.dev/)
1. Install [Docker](https://www.docker.com/)
1. `make up`
1. `make migrate`
1. `make install-sqlboiler`

## Generate code
1. `make generate-sqlboiler`
1. `make generate-mocks`

## Testing

1. Start the database by running `make up`
1. `make migrate`
1. `make test`
1. (Optional) After completion of your tests, `make down`

## Updating

1. `go get -u -d ./...`
1. `go mod tidy`
1. Update sqlboiler version in Makefile if it changed
1. `make up`
1. `make migrate`
1. `make install-sqlboiler`
1. `make generate-sqlboiler`

## Clean-up

To delete database and migrate again: `make clean`
