package controller

import (
	"net/http"
)

type IController interface {
	GetDeviceTelemetry(w http.ResponseWriter, r *http.Request)
	GetLatestDeviceTelemetry(w http.ResponseWriter, r *http.Request)
	Stop()
}
