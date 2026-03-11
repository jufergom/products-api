package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jufergom/products-api/internal/database"
)

func main() {
	uri := os.Getenv("MONGODB_URI")
	db, err := database.Connect(uri)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	coll := db.Collection("product")
	fmt.Println(coll)

	log.Println("Server starting")
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
