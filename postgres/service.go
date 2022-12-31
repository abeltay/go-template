package postgres

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Service represents a service for managing records in the database
type Service struct {
	db      boil.ContextExecutor
	beginTx func(ctx context.Context, opts *sql.TxOptions) (boil.ContextTransactor, error)
}

// NewService creates a new Service
func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
		beginTx: func(ctx context.Context, opts *sql.TxOptions) (boil.ContextTransactor, error) {
			return db.BeginTx(ctx, opts)
		},
	}
}
