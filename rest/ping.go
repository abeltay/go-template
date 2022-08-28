package rest

import (
	"fmt"
	"net/http"

	"github.com/abeltay/go-template/rest/logger"
	"go.uber.org/zap"
)

func (f Router) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := f.Service.Ping(ctx); err != nil {
		log := logger.ZapLogger(ctx)
		log.Error("error ping", zap.Error(err))
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "ok")
}
