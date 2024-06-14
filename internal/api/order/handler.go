package order

import (
	"context"
	"github.com/go-chi/chi/v5"
	"wb/internal/model"
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

func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/order/{id}", h.GetById)
	return r
}
