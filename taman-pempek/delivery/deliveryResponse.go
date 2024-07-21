package delivery

type DeliveryResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Whatsapp string `json:"whatsapp"`
}
