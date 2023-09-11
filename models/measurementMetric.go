package models

import (
	"encoding/json"
	"io"
)

type MeasurementMetric struct {
	Count  int    `json:"count"`
	Source string `json:"source"`
}

func NewMeasurementMetric(body io.ReadCloser) (*MeasurementMetric, error) {
	mm := &MeasurementMetric{}
	err := json.NewDecoder(body).Decode(mm)
	return mm, err
}
