package handlers

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/scfx/billing_microservice/billing"
	"github.com/scfx/billing_microservice/models"
)

type Measurement struct {
	log *log.Logger
}

func (m Measurement) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Read Measurment Metric from request body
	mm, err := models.NewMeasurementMetric(r.Body)
	if err != nil {
		m.log.Println("Error decoding measurement metric", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Get Tenant Id from Authorization Header
	// Split the string to get the base64-encoded part after "Basic "
	parts := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(parts) != 2 {
		m.log.Println("Error parsing authorization header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract the base64-encoded token
	base64Token := parts[1]

	decodedAuth, err := base64.StdEncoding.DecodeString(base64Token)
	if err != nil {
		m.log.Println("Error decoding authorization header", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	decodedString := string(decodedAuth)
	tenantId := strings.Split(decodedString, "/")[0]

	//Log Measurement and Authorization Header
	//Services
	var services []*billing.Metronome
	services = append(services, &billing.Metronome{})
	
	for _, service := range services {
		err := service.PublishMeasurement(mm, tenantId)
		if err != nil {
			m.log.Println("Error publishing measurement", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	//Log Measurment
	m.log.Printf("Measurement reported: %+v, at tenant: %s", mm, tenantId)
	//Return 201
	w.WriteHeader(http.StatusCreated)

}

func NewMeasurement(log *log.Logger) *Measurement {
	return &Measurement{log: log}
}
