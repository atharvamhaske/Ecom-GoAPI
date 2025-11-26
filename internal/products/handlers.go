package products

import (
	"log"
	"net/http"

	"github.com/atharvamhaske/Ecom-GoAPI/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Printf(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.Write(w, http.StatusOK, products)
}