package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jufergom/products-api/internal/repository"
)

type CustomerHandler struct {
	Repository *repository.CustomerRepository
}

func NewCustomerHandler(repository *repository.CustomerRepository) *CustomerHandler {
	return &CustomerHandler{Repository: repository}
}

func (h *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.Repository.FindAll()
	if err != nil {
		http.Error(w, "Error fetching customers", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	customer, err := h.Repository.FindByID(id)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
