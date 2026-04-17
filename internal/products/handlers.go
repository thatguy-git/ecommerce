package products

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) FindProductsbyId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var pgId pgtype.UUID
	if err := pgId.Scan(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	product, err := h.service.FindProductById(r.Context(), pgId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
