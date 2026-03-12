package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jufergom/products-api/internal/repository"
)

type ProductHandler struct {
	Repository *repository.ProductRepository
}

func NewHandler(repository *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{Repository: repository}
}

// Endpoint: GET /products
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Repository.FindAll()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

// Endpoint: GET /products/{id}
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, err := h.Repository.FindByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}
