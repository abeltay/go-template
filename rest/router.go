package rest

import (
	"fmt"
	"net/http"

	"github.com/abeltay/go-template/postgres"
	"github.com/abeltay/go-template/rest/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// Router represents the router
type Router struct {
	ZapLogger *zap.Logger
	Service   postgres.Postgres
}

// Handler returns a http.Handler to be run by a http.Server
func (f Router) Handler() http.Handler {
	r := chi.NewRouter()
	logger := logger.ZapLogMiddleware{
		Logger: f.ZapLogger,
	}
	r.Use(logger.Chain)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	r.Get("/ping", f.Ping)
	return r
}
