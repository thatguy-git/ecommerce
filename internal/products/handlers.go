package products

import (
	"log"
	"net/http"

	"github.com/thatguy-git/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	val, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, val)
}
