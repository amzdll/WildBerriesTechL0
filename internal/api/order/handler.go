package order

import (
	"context"
	"wb/internal/model"

	"github.com/go-chi/chi/v5"
)

type Service interface {
	Get(ctx context.Context, id string) (model.Order, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() (string, *chi.Mux) {
	route := "/order"
	r := chi.NewRouter()

	r.Get("/{id}", h.GetById)

	return route, r
}
