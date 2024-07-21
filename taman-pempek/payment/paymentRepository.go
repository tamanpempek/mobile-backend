package payment

import (
	"errors"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindAll() ([]Payment, error)
	FindPaymentByID(ID int) (Payment, error)
	FindPaymentByUserAndStatus(userID int, paymentStatus string) ([]Payment, error)
	FindPaymentByStatus(paymentStatus string) ([]Payment, error)
	CreatePayment(payment Payment) (Payment, error)
	UpdatePayment(payment Payment) (Payment, error)
	DeletePayment(payment Payment) (Payment, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindAll() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error
	return payments, err
}

func (r *repository) FindPaymentByID(ID int) (Payment, error) {
	var payment Payment
	err := r.db.First(&payment, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Payment{}, errors.New("Payment not found")
	}
	return payment, err
}

func (r *repository) FindPaymentByUserAndStatus(userID int, paymentStatus string) ([]Payment, error) {
	var payments []Payment
	err := r.db.Where("user_id = ? AND payment_status = ?", userID, paymentStatus).Find(&payments).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []Payment{}, errors.New("Payment not found")
	}
	return payments, err
}

func (r *repository) FindPaymentByStatus(paymentStatus string) ([]Payment, error) {
	var payments []Payment
	err := r.db.Where("payment_status = ?", paymentStatus).Find(&payments).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []Payment{}, errors.New("Payment not found")
	}
	return payments, err
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePayment(payment Payment) (Payment, error) {
	err := r.db.Create(&payment).Error
	return payment, err
}

func (r *repository) UpdatePayment(payment Payment) (Payment, error) {
	err := r.db.Save(&payment).Error
	return payment, err
}

func (r *repository) DeletePayment(payment Payment) (Payment, error) {
	err := r.db.Delete(&payment).Error
	return payment, err
}
