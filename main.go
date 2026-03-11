package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting")
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
