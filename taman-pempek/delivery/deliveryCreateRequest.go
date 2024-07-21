package delivery

type DeliveryCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Whatsapp string `json:"whatsapp" binding:"required"`
}
