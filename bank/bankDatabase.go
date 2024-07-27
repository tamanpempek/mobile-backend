package bank

import "time"

type Bank struct {
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    int        `gorm:"column:user_id;type:varchar(255)"`
	Type      string     `gorm:"column:type;type:varchar(255)"`
	Name      string     `gorm:"column:name;type:varchar(255)"`
	Number    string     `gorm:"column:number;type:varchar(255)"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
