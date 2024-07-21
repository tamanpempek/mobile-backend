package cart

import "encoding/json"

type CartUpdateRequest struct {
	ProductID  json.Number `json:"product_id,omitempty"`
	PaymentID  json.Number `json:"payment_id,omitempty"`
	Quantity   json.Number `json:"quantity,omitempty"`
	TotalPrice json.Number `json:"total_price,omitempty"`
	IsActived  string      `json:"isActived,omitempty"`
}
