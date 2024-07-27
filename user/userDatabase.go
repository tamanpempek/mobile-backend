package user

import "time"

type User struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name;type:varchar(255)"`
	Email     string    `gorm:"column:email;type:varchar(255);unique"`
	Password  string    `gorm:"column:password;type:varchar(255)"`
	Whatsapp  string    `gorm:"column:whatsapp;type:varchar(255)"`
	Gender    string    `gorm:"column:gender;type:varchar(255)"`
	Role      string    `gorm:"column:role;type:varchar(255);"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
