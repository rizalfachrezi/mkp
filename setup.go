package main

import (
	// "github.com/gorilla/mux" ERORRR
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
