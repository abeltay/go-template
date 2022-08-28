package postgres

import "context"

// Postgres holds the methods for the package
type Postgres interface {
	Ping(ctx context.Context) error
}
