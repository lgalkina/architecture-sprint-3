package entities

type TelemetryData struct {
	ID          string  `json:"id"`
	DeviceID    string  `json:"device_id"`
	Temperature float64 `json:"temperature"`
	Timestamp   string  `json:"timestamp"`
}
