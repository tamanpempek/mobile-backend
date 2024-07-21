package delivery

import (
	"errors"

	"gorm.io/gorm"
)

type DeliveryService interface {
	FindAll() ([]Delivery, error)
	FindDeliveryByID(ID int) (Delivery, error)
	CreateDelivery(delivery DeliveryCreateRequest) (Delivery, error)
	UpdateDelivery(ID int, delivery DeliveryUpdateRequest) (Delivery, error)
	DeleteDelivery(ID int) (Delivery, error)
}

type service struct {
	deliveryRepository DeliveryRepository
}

func NewService(deliveryRepository DeliveryRepository) *service {
	return &service{deliveryRepository}
}

func (s *service) FindAll() ([]Delivery, error) {
	return s.deliveryRepository.FindAll()
}

func (s *service) FindDeliveryByID(ID int) (Delivery, error) {
	return s.deliveryRepository.FindDeliveryByID(ID)
}

func (s *service) CreateDelivery(deliveryRequest DeliveryCreateRequest) (Delivery, error) {
	deliveryData := Delivery{
		Name:     deliveryRequest.Name,
		Whatsapp: deliveryRequest.Whatsapp,
	}

	delivery, err := s.deliveryRepository.CreateDelivery(deliveryData)

	return delivery, err
}

func (s *service) UpdateDelivery(ID int, deliveryRequest DeliveryUpdateRequest) (Delivery, error) {
	delivery, err := s.deliveryRepository.FindDeliveryByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Delivery{}, errors.New("Delivery not found")
	}

	if err != nil {
		return Delivery{}, err
	}

	if deliveryRequest.Name != "" {
		delivery.Name = deliveryRequest.Name
	}
	if deliveryRequest.Whatsapp != "" {
		delivery.Whatsapp = deliveryRequest.Whatsapp
	}

	return s.deliveryRepository.UpdateDelivery(delivery)
}

func (s *service) DeleteDelivery(ID int) (Delivery, error) {
	delivery, err := s.deliveryRepository.FindDeliveryByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Delivery{}, errors.New("Delivery not found")
	}

	if err != nil {
		return Delivery{}, err
	}

	return s.deliveryRepository.DeleteDelivery(delivery)
}
