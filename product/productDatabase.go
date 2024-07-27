package product

import "time"

type Product struct {
	ID          uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID      int       `gorm:"column:user_id;type:varchar(255)"`
	CategoryID  int       `gorm:"column:category_id;type:varchar(255)"`
	Name        string    `gorm:"column:name;type:varchar(255)"`
	Image       string    `gorm:"column:image;type:varchar(255)"`
	Description string    `gorm:"column:description;type:text"`
	Price       int       `gorm:"column:price"`
	Stock       int       `gorm:"column:stock"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
