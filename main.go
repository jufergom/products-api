package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jufergom/products-api/internal/database"
	"github.com/jufergom/products-api/internal/handler"
	"github.com/jufergom/products-api/internal/repository"
)

func main() {
	uri := os.Getenv("MONGODB_URI")
	db, err := database.Connect(uri)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	productRepository := repository.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepository)

	customerRepository := repository.NewCustomerRepository(db)
	customerHandler := handler.NewCustomerHandler(customerRepository)

	log.Println("Server starting")
	r := mux.NewRouter()
	r.HandleFunc("/api/products", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", productHandler.GetProductByID).Methods("GET")
	r.HandleFunc("/api/customers", customerHandler.GetAllCustomers).Methods("GET")
	r.HandleFunc("/api/customers/{id}", customerHandler.GetCustomerByID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
