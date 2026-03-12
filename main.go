package main

import (
	"fmt"
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

	productRepository := repository.NewRepository(db)
	productHandler := handler.NewHandler(productRepository)

	log.Println("Server starting")
	r := mux.NewRouter()
	r.HandleFunc("/api/products", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", productHandler.GetProductByID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
