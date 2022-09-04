package main

import (
	"fmt"
	"log"

	"github.com/abeltay/go-template/env"
	"github.com/abeltay/go-template/postgres"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	options, err := env.LoadOSEnv()
	if err != nil {
		log.Fatalln("Error loading environment settings, exiting the program: ", err)
	}
	db, err := postgres.OpenDB(options)
	if err != nil {
		log.Fatalln("Database error: ", err)
	}
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalln("Migrate error: ", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
