package entities

type Module struct {
	ID       int64  `json:"id"`
	DeviceID int64  `json:"device_id"`
	Name     string `json:"name"`
}
