package delivery

import (
	"errors"

	"gorm.io/gorm"
)

type DeliveryRepository interface {
	FindAll() ([]Delivery, error)
	FindDeliveryByID(ID int) (Delivery, error)
	CreateDelivery(delivery Delivery) (Delivery, error)
	UpdateDelivery(delivery Delivery) (Delivery, error)
	DeleteDelivery(delivery Delivery) (Delivery, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Delivery, error) {
	var deliveries []Delivery
	err := r.db.Find(&deliveries).Error
	return deliveries, err
}

func (r *repository) FindDeliveryByID(ID int) (Delivery, error) {
	var delivery Delivery
	err := r.db.First(&delivery, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Delivery{}, errors.New("Delivery not found")
	}
	return delivery, err
}

func (r *repository) CreateDelivery(delivery Delivery) (Delivery, error) {
	err := r.db.Create(&delivery).Error
	return delivery, err
}

func (r *repository) UpdateDelivery(delivery Delivery) (Delivery, error) {
	err := r.db.Save(&delivery).Error
	return delivery, err
}

func (r *repository) DeleteDelivery(delivery Delivery) (Delivery, error) {
	err := r.db.Delete(&delivery).Error
	return delivery, err
}
