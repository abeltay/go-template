package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abeltay/go-template/rest/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (f Router) AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.ZapLogger(ctx)
	type input struct {
		Name string `json:"name"`
	}
	var i input
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		log.Info("param error", zap.Error(err))
		http.Error(w, "error", http.StatusBadRequest)
		return
	}
	id, err := f.Service.AddUser(ctx, i.Name)
	if err != nil {
		log.Info("service error", zap.Error(err))
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User created with id: ", id)
}

func (f Router) UserFullName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log := logger.ZapLogger(ctx)
		log.Info("param error", zap.Error(err))
		http.Error(w, "error", http.StatusBadRequest)
		return
	}
	name, err := f.Service.UserFullName(ctx, id)
	if err != nil {
		log := logger.ZapLogger(ctx)
		log.Info("retrieve error", zap.Error(err))
		http.Error(w, "error", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "User's name is: ", name)
}
