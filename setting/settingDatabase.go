package setting

import "time"

type Setting struct {
	ID          uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	Image       string     `gorm:"column:image;type:varchar(255)"`
	Description string     `gorm:"column:description;type:varchar(255)"`
	Contact     string     `gorm:"column:contact;type:varchar(255)"`
	Email       string     `gorm:"column:email;type:varchar(255)"`
	Instagram   string     `gorm:"column:instagram;type:varchar(255)"`
	Website     string     `gorm:"column:website;type:varchar(255)"`
	CreatedAt   *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
