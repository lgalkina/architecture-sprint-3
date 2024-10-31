package entities

type House struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Address string `json:"address"`
}
