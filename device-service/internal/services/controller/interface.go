package controller

import "net/http"

type IController interface {
	GetDeviceInfo(w http.ResponseWriter, r *http.Request)
	UpdateDeviceStatus(w http.ResponseWriter, r *http.Request)
	SendDeviceCommand(w http.ResponseWriter, r *http.Request)
	Stop()
}
