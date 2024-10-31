package entities

type TelemetryData struct {
	ID          int64   `json:"id"`
	DeviceID    int64   `json:"device_id"`
	Temperature float64 `json:"temperature"`
	Timestamp   string  `json:"timestamp"`
}
