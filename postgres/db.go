package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/abeltay/go-template/env"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// OpenDB returns a new instance of DB associated with the given datasource name.
// caller should "defer db.Close()" after checking for error
func OpenDB(cfg env.Config) (*sql.DB, error) {
	// Connect to the database.
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseName)
	config, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	db := stdlib.OpenDB(*config)
	db.SetMaxOpenConns(cfg.DatabaseMaxConn)
	db.SetMaxIdleConns(cfg.DatabaseMaxConn)
	db.SetConnMaxLifetime(time.Duration(cfg.DatabaseConnMaxLifetime) * time.Second)
	return db, nil
}
