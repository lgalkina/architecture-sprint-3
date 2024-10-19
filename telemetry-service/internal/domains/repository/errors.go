package repository

import (
	"fmt"
)

type DeviceTelemetryNotFoundError struct {
	ID string
}

func NewDeviceTelemetryNotFoundError(id string) error {
	return &DeviceTelemetryNotFoundError{id}
}

func (e *DeviceTelemetryNotFoundError) Error() string {
	return fmt.Sprintf("telemetry for device device with ID %s not found", e.ID)
}
