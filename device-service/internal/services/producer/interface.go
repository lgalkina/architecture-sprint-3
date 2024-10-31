package producer

import "device-service/internal/domains/entities"

type ICommandProducer interface {
	SendCommand(entities.DeviceCommand) error
	Close()
}
