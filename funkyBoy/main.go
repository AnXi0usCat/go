package main

import (
	"github.com/gorilla/mux"
	"funkyBoy/handlers"
	"net/http"
	"log"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/scale/{value}", handlers.GetScaledValue).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
