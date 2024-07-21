package payment

import (
	"mime/multipart"
)

type PaymentCreateRequest struct {
	UserID        int                  `form:"user_id" binding:"required"`
	DeliveryID    int                  `form:"delivery_id"`
	TotalPrice    int                  `form:"total_price" binding:"required"`
	Image         multipart.FileHeader `form:"image" binding:"required"`
	Address       string               `form:"address" binding:"required"`
	Whatsapp      string               `form:"whatsapp" binding:"required"`
	PaymentStatus string               `form:"payment_status" binding:"required"`
	DeliveryName  string               `form:"delivery_name" binding:"required"`
	Resi          string               `form:"resi" binding:"required"`
}
