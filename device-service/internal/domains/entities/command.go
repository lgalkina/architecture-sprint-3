package entities

type DeviceCommand struct {
	Command string  `json:"command"`
	Value   float64 `json:"value"`
}

type DeviceCommandResponse struct {
	Command string  `json:"command"`
	Value   float64 `json:"value"`
	Status  string  `json:"status"`
}
