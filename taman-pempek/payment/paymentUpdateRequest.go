package payment

type PaymentUpdateRequest struct {
	DeliveryID    int    `form:"delivery_id,omitempty"`
	TotalPrice    int    `form:"total_price,omitempty"`
	Address       string `form:"address,omitempty"`
	Whatsapp      string `form:"whatsapp,omitempty"`
	PaymentStatus string `form:"payment_status,omitempty"`
	DeliveryName  string `form:"delivery_name,omitempty"`
	Resi          string `form:"resi,omitempty"`
}
