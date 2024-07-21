package cart

import (
	"encoding/json"
	"time"
)

type Cart struct {
	ID         uint64      `gorm:"column:id;primaryKey;autoIncrement"`
	UserID     int         `gorm:"column:user_id;type:varchar(255)"`
	ProductID  json.Number `gorm:"column:product_id;type:varchar(255)"`
	PaymentID  json.Number `gorm:"column:payment_id;type:varchar(255)"`
	Quantity   json.Number `gorm:"column:quantity;type:varchar(255)"`
	TotalPrice json.Number `gorm:"column:total_price;type:varchar(255)"`
	IsActived  string      `gorm:"column:isActived;type:varchar(255)"`
	CreatedAt  time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"column:updated_at;autoUpdateTime"`
}
