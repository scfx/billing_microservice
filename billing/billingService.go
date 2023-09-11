package billing

import "github.com/scfx/billing_microservice/models"

type billingService interface {
	PublishMeasurement(measurementMetric *models.MeasurementMetric, tenantId string) error
}
