package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"telemetry-service/internal/domains/repository"
	"telemetry-service/internal/services/consumer"
)

type controller struct {
	telemetryRepo repository.ITelemetryRepository
	consumer      consumer.IConsumerService
}

func NewController() IController {
	return &controller{
		telemetryRepo: repository.NewTelemetryRepository(),
		consumer:      consumer.NewConsumer(),
	}
}

func (c *controller) GetDeviceTelemetry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceID := vars["device_id"]
	telemetry, err := c.telemetryRepo.GetDeviceTelemetry(deviceID)
	if err != nil {
		if _, ok := err.(*repository.DeviceTelemetryNotFoundError); ok {
			http.Error(w, "Telemetry device not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error getting device: %v", err), http.StatusInternalServerError)
		}
		return
	}
	if len(telemetry) == 0 {
		http.Error(w, "Telemetry device not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(telemetry)
}

func (c *controller) GetLatestDeviceTelemetry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceID := vars["device_id"]
	telemetry, err := c.telemetryRepo.GetLatestDeviceTelemetry(deviceID)
	if err != nil {
		if _, ok := err.(*repository.DeviceTelemetryNotFoundError); ok {
			http.Error(w, "Telemetry device not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error getting device: %v", err), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(telemetry)
}

func (c *controller) Stop() {
	c.consumer.Close()
}
