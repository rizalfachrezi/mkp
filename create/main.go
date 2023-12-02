package main

import (
	// "github.com/gorilla/mux" ERROR
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/terminal/create", createTerminal).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
