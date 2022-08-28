package main

import (
	"log"
	"time"

	"github.com/abeltay/go-kit/web"
	"github.com/abeltay/go-template/env"
	"github.com/abeltay/go-template/postgres"
	"github.com/abeltay/go-template/rest"
	"go.uber.org/zap"
)

func main() {
	options, err := env.LoadOSEnv()
	if err != nil {
		log.Fatalln("Error loading environment settings, exiting the program: ", err)
	}

	var config zap.Config
	if options.Production {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}
	config.DisableCaller = true

	zaplogger, err := config.Build()
	if err != nil {
		log.Fatalln("Error setting up logger, exiting: ", err)
	}
	defer zaplogger.Sync()

	db, err := postgres.OpenDB(options)
	if err != nil {
		zaplogger.Fatal("Database", zap.Error(err))
	}
	defer db.Close()

	r := rest.Router{
		ZapLogger: zaplogger,
		Service:   postgres.NewService(db),
	}

	// Create a new server and set timeout values.
	// Docker has a default timeout of 10 seconds https://docs.docker.com/compose/reference/stop/
	// Create a context to attempt a graceful 8 second shutdown.
	server := web.Server{
		Addr:    ":8000",
		Handler: r.Handler(),
		Timeout: 8 * time.Second,
	}

	zaplogger.Info("Starting server", zap.Any("env", options))
	zaplogger.Info("ListenAndServe", zap.Error(server.ListenAndServe()))
}
