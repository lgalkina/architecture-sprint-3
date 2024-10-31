package repository

import (
	"fmt"
)

type DeviceNotFoundError struct {
	ID string
}

func NewDeviceNotFoundError(id string) error {
	return &DeviceNotFoundError{id}
}

func (e *DeviceNotFoundError) Error() string {
	return fmt.Sprintf("device with ID %s not found", e.ID)
}
