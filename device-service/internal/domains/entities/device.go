package entities

type Device struct {
	ID           string     `json:"id"`
	TypeID       string     `json:"type_id"`
	HouseID      string     `json:"house_id"`
	SerialNumber string     `json:"serial_number"`
	Status       string     `json:"status"`
	DeviceType   DeviceType `json:"device_type"`
	Modules      []Module   `json:"modules"`
	House        House      `json:"house"`
}
