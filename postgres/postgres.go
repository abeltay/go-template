package postgres

import "context"

// Postgres holds the methods for the package
type Postgres interface {
	Ping(ctx context.Context) error
	AddUser(ctx context.Context, fullName string) (int, error)
	UserFullName(ctx context.Context, id int) (string, error)
}
