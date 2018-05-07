package main

import (
	"github.com/gorilla/mux"
	"funkyBoy/handlers"
	"net/http"
	"log"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/scale/{value}", handlers.GetScaledValueTest).Methods("GET")
	router.HandleFunc("/scale", handlers.GetScaledValue).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
