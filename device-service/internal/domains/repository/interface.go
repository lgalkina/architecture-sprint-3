package repository

import (
	"device-service/internal/domains/entities"
)

type IDeviceRepository interface {
	GetDeviceInfo(id string) (*entities.Device, error)
	UpdateDeviceStatus(id string, status string) error
}
