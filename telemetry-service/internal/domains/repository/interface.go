package repository

import (
	"telemetry-service/internal/domains/entities"
)

type ITelemetryRepository interface {
	GetDeviceTelemetry(id string) ([]entities.TelemetryData, error)
	GetLatestDeviceTelemetry(id string) (*entities.TelemetryData, error)
	SaveDeviceTelemetry(data *entities.TelemetryData) error
}
