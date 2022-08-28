package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/abeltay/go-template/env"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

// OpenDB returns a new instance of DB associated with the given datasource name.
// caller should "defer db.Close()" after checking for error
func OpenDB(options env.Options) (*sql.DB, error) {
	// Connect to the database.
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", options.DatabaseUser, options.DatabasePassword, options.DatabaseHost, options.DatabasePort, options.DatabaseName)
	config, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	db := stdlib.OpenDB(*config)
	db.SetMaxOpenConns(options.DatabaseMaxConn)
	db.SetMaxIdleConns(options.DatabaseMaxConn)
	db.SetConnMaxLifetime(time.Duration(options.DatabaseConnMaxLifetime) * time.Second)
	return db, nil
}
