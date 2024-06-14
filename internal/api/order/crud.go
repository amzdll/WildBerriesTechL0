package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"wb/internal/api/order/response"
)

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	order, err := h.service.Get(r.Context(), id)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response.OrderFromModel(order))
}
