package rest

import (
	"fmt"
	"net/http"

	"github.com/abeltay/go-template/postgres"
	"github.com/abeltay/go-template/rest/logger"
	"go.uber.org/zap"
)

// Router represents the router
type Router struct {
	ZapLogger *zap.Logger
	Service   postgres.Postgres
}

// Handler returns a http.Handler to be run by a http.Server
func (f Router) Handler() http.Handler {
	mux := http.NewServeMux()
	logger := logger.ZapLogMiddleware{
		Logger: f.ZapLogger,
	}
	mux.HandleFunc("GET /livez", logger.Chain(liveliness))
	mux.HandleFunc("GET /readyz", logger.Chain(f.Ping))

	mux.HandleFunc("POST /user/", logger.Chain(f.AddUser))
	mux.HandleFunc("GET /user/{id}", logger.Chain(f.UserFullName))

	return mux
}

func liveliness(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
