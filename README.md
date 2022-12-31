# Go Template

## Setup

1. Install [Go](https://go.dev/)
1. Install [Docker](https://www.docker.com/)
1. Install SQL migrate: `go install github.com/rubenv/sql-migrate/...@latest`
1. `make up`
1. `make seed`

## Testing

1. Start the database by running `make up`
1. `make test`
1. (Optional) After completion of your tests, `make down`

## Clean-up

1. `make clean`
