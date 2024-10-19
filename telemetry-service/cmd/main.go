package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"telemetry-service/internal/services/controller"
)

func main() {
	setup()

	controller := controller.NewController()
	defer controller.Stop()

	r := mux.NewRouter()
	r.HandleFunc("/devices/{device_id}/telemetry", controller.GetDeviceTelemetry).Methods("GET")
	r.HandleFunc(" /devices/{device_id}/telemetry/latest", controller.GetLatestDeviceTelemetry).Methods("GET")

	log.Println("Starting server on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

func setup() {
	if postgresURL := os.Getenv("TELEMETRY_POSTGRES_URL"); postgresURL == "" {
		os.Setenv("TELEMETRY_POSTGRES_URL", "postgres://postgres:postgres@localhost:5433/telemetry_db?sslmode=disable")
	}
	if kafkaServers := os.Getenv("KAFKA_SERVERS"); kafkaServers == "" {
		os.Setenv("KAFKA_SERVERS", "localhost:9092")
	}
}
