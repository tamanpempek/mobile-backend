package payment

import (
	"errors"

	"gorm.io/gorm"
)

type PaymentService interface {
	FindAll() ([]Payment, error)
	FindPaymentByID(ID int) (Payment, error)
	FindPaymentByUserAndStatus(userID int, paymentStatus string) ([]Payment, error)
	FindPaymentByStatus(paymentStatus string) ([]Payment, error)
	CreatePayment(payment PaymentCreateRequest) (Payment, error)
	UpdatePayment(ID int, payment PaymentUpdateRequest) (Payment, error)
	DeletePayment(ID int) (Payment, error)
}

type service struct {
	paymentRepository PaymentRepository
}

func NewService(paymentRepository PaymentRepository) *service {
	return &service{paymentRepository}
}

func (s *service) FindAll() ([]Payment, error) {
	return s.paymentRepository.FindAll()
}

func (s *service) FindPaymentByID(ID int) (Payment, error) {
	return s.paymentRepository.FindPaymentByID(ID)
}

func (s *service) FindPaymentByUserAndStatus(userID int, paymentStatus string) ([]Payment, error) {
	return s.paymentRepository.FindPaymentByUserAndStatus(userID, paymentStatus)
}

func (s *service) FindPaymentByStatus(paymentStatus string) ([]Payment, error) {
	return s.paymentRepository.FindPaymentByStatus(paymentStatus)
}

func (s *service) CreatePayment(paymentRequest PaymentCreateRequest) (Payment, error) {
	paymentData := Payment{
		UserID:        paymentRequest.UserID,
		DeliveryID:    paymentRequest.DeliveryID,
		TotalPrice:    paymentRequest.TotalPrice,
		Image:         paymentRequest.Image.Filename,
		Address:       paymentRequest.Address,
		Whatsapp:      paymentRequest.Whatsapp,
		PaymentStatus: paymentRequest.PaymentStatus,
		DeliveryName:  paymentRequest.DeliveryName,
		Resi:          paymentRequest.Resi,
	}

	payment, err := s.paymentRepository.CreatePayment(paymentData)

	return payment, err
}

func (s *service) UpdatePayment(ID int, paymentRequest PaymentUpdateRequest) (Payment, error) {
	payment, err := s.paymentRepository.FindPaymentByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Payment{}, errors.New("Payment not found")
	}

	if err != nil {
		return Payment{}, err
	}
	if paymentRequest.DeliveryID != 0 {
		payment.DeliveryID = paymentRequest.DeliveryID
	}
	if paymentRequest.TotalPrice != 0 {
		payment.TotalPrice = paymentRequest.TotalPrice
	}
	if paymentRequest.Address != "" {
		payment.Address = paymentRequest.Address
	}
	if paymentRequest.Whatsapp != "" {
		payment.Whatsapp = paymentRequest.Whatsapp
	}
	if paymentRequest.PaymentStatus != "" {
		payment.PaymentStatus = paymentRequest.PaymentStatus
	}
	if paymentRequest.DeliveryName != "" {
		payment.DeliveryName = paymentRequest.DeliveryName
	}
	if paymentRequest.Resi != "" {
		payment.Resi = paymentRequest.Resi
	}

	return s.paymentRepository.UpdatePayment(payment)
}

func (s *service) DeletePayment(ID int) (Payment, error) {
	payment, err := s.paymentRepository.FindPaymentByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Payment{}, errors.New("Payment not found")
	}

	if err != nil {
		return Payment{}, err
	}

	return s.paymentRepository.DeletePayment(payment)
}
