package main

import (
	"log"
	"net/http"
	"os"

	"device-service/internal/services/controller"
	"github.com/gorilla/mux"
)

func main() {
	setup()

	controller := controller.NewController()
	defer controller.Stop()

	r := mux.NewRouter()
	r.HandleFunc("/devices/{device_id}", controller.GetDeviceInfo).Methods("GET")
	r.HandleFunc("/devices/{device_id}/status", controller.UpdateDeviceStatus).Methods("PUT")
	r.HandleFunc("/devices/{device_id}/commands", controller.SendDeviceCommand).Methods("POST")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setup() {
	if postgresURL := os.Getenv("DEVICE_POSTGRES_URL"); postgresURL == "" {
		os.Setenv("DEVICE_POSTGRES_URL", "postgres://postgres:postgres@localhost:5432/device_db?sslmode=disable")
	}
	if kafkaServers := os.Getenv("KAFKA_SERVERS"); kafkaServers == "" {
		os.Setenv("KAFKA_SERVERS", "localhost:9092")
	}
}
