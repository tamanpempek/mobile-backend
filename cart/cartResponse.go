package cart

import "encoding/json"

type CartResponse struct {
	ID         uint64      `json:"id"`
	UserID     int         `json:"user_id"`
	ProductID  json.Number `json:"product_id"`
	PaymentID  json.Number `json:"payment_id"`
	Quantity   json.Number `json:"quantity"`
	TotalPrice json.Number `json:"total_price"`
	IsActived  string      `json:"isActived"`
}
