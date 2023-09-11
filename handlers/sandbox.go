package handlers

import (
	"log"
	"net/http"

	"github.com/scfx/billing_microservice/billing"
	"github.com/scfx/billing_microservice/models"
)

type SandBox struct {
	log *log.Logger
}

// New Test handler
func NewSandBox(log *log.Logger) *SandBox {

	return &SandBox{log: log}
}

func (s SandBox) Measurement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Read Measurment Metric from request body
		mm, err := models.NewMeasurementMetric(r.Body)
		if err != nil {
			log.Println("Error decoding measurement metric", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//Services
		var services []*billing.Metronome
		services = append(services, &billing.Metronome{})
		//Dummy tenant Id
		tenantId := "123456789"
		//Publish Measurement
		for _, service := range services {
			err := service.PublishMeasurement(mm, tenantId)
			if err != nil {
				s.log.Println("Error publishing measurement", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		//Log Measurment
		s.log.Printf("Measurement reported: %+v, at tenant: %s", mm, tenantId)
		//Return 201
		w.WriteHeader(http.StatusCreated)
	}
}
