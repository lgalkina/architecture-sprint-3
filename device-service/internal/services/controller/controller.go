package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"device-service/internal/domains/entities"
	"device-service/internal/domains/repository"
	"device-service/internal/services/producer"
	"github.com/gorilla/mux"
)

type controller struct {
	deviceRepo      repository.IDeviceRepository
	commandProducer producer.ICommandProducer
}

func NewController() IController {
	return &controller{
		deviceRepo:      repository.NewDeviceRepository(),
		commandProducer: producer.NewCommandProducer(),
	}
}

func (c *controller) GetDeviceInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceID := vars["device_id"]
	device, err := c.deviceRepo.GetDeviceInfo(deviceID)
	if err != nil {
		if _, ok := err.(*repository.DeviceNotFoundError); ok {
			http.Error(w, "Device not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error getting device: %v", err), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(device)
}

func (c *controller) UpdateDeviceStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceID := vars["device_id"]
	var status entities.DeviceStatus
	_ = json.NewDecoder(r.Body).Decode(&status)
	err := c.deviceRepo.UpdateDeviceStatus(deviceID, status.Status)
	if err != nil {
		if _, ok := err.(*repository.DeviceNotFoundError); ok {
			http.Error(w, "Device not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error updating device: %v", err), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func (c *controller) SendDeviceCommand(w http.ResponseWriter, r *http.Request) {
	var command entities.DeviceCommand
	_ = json.NewDecoder(r.Body).Decode(&command)
	if err := c.commandProducer.SendCommand(command); err != nil {
		http.Error(w, fmt.Sprintf("Error sending device command: %v", err), http.StatusInternalServerError)
		return
	}
	// Логика для отправки команды устройству
	response := entities.DeviceCommandResponse{
		Command: command.Command,
		Value:   command.Value,
		Status:  "sent",
	}
	json.NewEncoder(w).Encode(response)
}

func (c *controller) Stop() {
	c.commandProducer.Close()
}
