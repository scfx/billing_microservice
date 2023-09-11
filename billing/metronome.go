package billing

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/scfx/billing_microservice/models"
)

type Metronome struct {
}

type MetronomeMeasurement struct {
	TransactionId string                         `json:"transaction_id"`
	CustomerId    string                         `json:"customer_id"`
	Timestamp     string                         `json:"timestamp"`
	EventType     string                         `json:"event_type"`
	Properties    MetronomeMeasurementProperties `json:"properties"`
}

type MetronomeMeasurementProperties struct {
	Count  int    `json:"count"`
	Source string `json:"source"`
}

func (m Metronome) PublishMeasurement(measurementMetric *models.MeasurementMetric, tenantId string) error {
	//Create MetronomeMeasurementMetric
	metrics := make([]*MetronomeMeasurement, 1)
	metrics[0] = &MetronomeMeasurement{
		TransactionId: uuid.New().String(),
		CustomerId:    tenantId,
		Timestamp:     time.Now().Format(time.RFC3339),
		EventType:     "measurement",
		Properties: MetronomeMeasurementProperties{
			Count:  measurementMetric.Count,
			Source: measurementMetric.Source,
		},
	}

	// Marshal MetronomeMeasurementMetric to JSON
	json, err := json.Marshal(metrics)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", m.baseUrl()+"/ingest", bytes.NewBuffer(json))
	//Set Headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.token())

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return errors.New("Publish to Metronome failed with status code: " + string(resp.StatusCode) + " and message: " + resp.Status)
	}
	// Publish JSON to Metronome
	return nil

}

func (m Metronome) baseUrl() string {
	return "https://api.metronome.com/v1"
}

func (m Metronome) token() string {
	return "1cb46242edf5057eda4bbbcdbd2edd65b9a1bf5a54a921866cb1556140e73350"
}
