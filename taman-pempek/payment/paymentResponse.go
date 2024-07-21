package payment

type PaymentResponse struct {
	ID            uint64 `json:"id"`
	UserID        int    `json:"user_id"`
	DeliveryID    int    `json:"delivery_id"`
	TotalPrice    int    `json:"total_price"`
	Image         string `json:"image"`
	Address       string `json:"address"`
	Whatsapp      string `json:"whatsapp"`
	PaymentStatus string `json:"payment_status"`
	DeliveryName  string `json:"delivery_name"`
	Resi          string `json:"resi"`
}
