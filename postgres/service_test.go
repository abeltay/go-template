package postgres

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/abeltay/go-template/env"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var db *sql.DB

func TestMain(m *testing.M) {
	options, err := env.LoadOSEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err = OpenDB(options)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	os.Exit(m.Run())
}

func newFakeService(tx boil.ContextExecutor) *Service {
	return &Service{
		db: tx,
		beginTx: func(ctx context.Context, opts *sql.TxOptions) (boil.ContextTransactor, error) {
			return fakeTx{tx}, nil
		},
	}
}

type fakeTx struct {
	boil.ContextExecutor
}

func (fakeTx) Commit() error   { return nil }
func (fakeTx) Rollback() error { return nil }

func TestNewService(t *testing.T) {
	NewService(db)
}
