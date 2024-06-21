package order

import (
	"net/http"
	"wb/internal/api/order/response"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// GetById fetches an order by ID.
//
// @Summary Fetches an order by its ID.
// @Description Retrieves an order identified by the provided ID.
// @Tags Orders
// @Param id path string true "Order ID"
// @Accept json
// @Produce json
// @Success 200 {object} response.Order "Order successfully get"
// @Failure 404 {string} Not found
// @Router /order/{id} [get]
func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	order, err := h.service.Get(r.Context(), id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response.OrderFromModel(order))
}
