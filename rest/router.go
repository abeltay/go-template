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
	r.Get("/livez", liveliness)
	r.Get("/readyz", f.Ping)

	r.Group(f.group)
	return r
}

func liveliness(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func (f Router) group(r chi.Router) {
	r.Route("/user", f.users)
}

func (f Router) users(r chi.Router) {
	r.Post("/", f.AddUser)
	r.Get("/{id}", f.UserFullName)
}
