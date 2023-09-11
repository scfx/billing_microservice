package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/scfx/billing_microservice/handlers"
)

func main() {
	//Create logger
	log := log.New(os.Stdout, "Billing", log.LstdFlags|log.Lshortfile)

	r := mux.NewRouter()

	//Create handlers
	hh := handlers.NewHealth(log)
	mh := handlers.NewMeasurement(log)
	sh := handlers.NewSandBox(log)

	//Sandbox Endpoint
	r.HandleFunc("/sandbox/measurement", sh.Measurement()).Methods("POST")
	//Health Endpoint
	r.Handle("/health", hh).Methods("GET")
	//Measurement Endpoint
	r.Handle("/measurement", mh).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":80", r)
}
