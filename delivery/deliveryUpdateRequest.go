package delivery

type DeliveryUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Whatsapp string `json:"whatsapp,omitempty"`
}
