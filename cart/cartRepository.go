package cart

import (
	"errors"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindAll() ([]Cart, error)
	FindCartByID(ID int) (Cart, error)
	FindCartsByPaymentID(paymentID int) ([]Cart, error)
	FindCartsByProductID(productID int) ([]Cart, error)
	FindStatusCardByUser(userID int, isActived string) ([]Cart, error)
	SumTotalPriceByUser(userID int, isActived string) (int, error)
	CreateCart(cart Cart) (Cart, error)
	UpdateCart(cart Cart) (Cart, error)
	DeleteCart(cart Cart) (Cart, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindAll() ([]Cart, error) {
	var carts []Cart
	err := r.db.Find(&carts).Error
	return carts, err
}

func (r *repository) FindCartByID(ID int) (Cart, error) {
	var cart Cart
	err := r.db.First(&cart, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Cart{}, errors.New("Cart not found")
	}
	return cart, err
}

func (r *repository) FindCartsByPaymentID(paymentID int) ([]Cart, error) {
	var carts []Cart
	err := r.db.Where("payment_id = ?", paymentID).Find(&carts).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []Cart{}, errors.New("Cart not found")
	}
	return carts, err
}

func (r *repository) FindCartsByProductID(productID int) ([]Cart, error) {
	var carts []Cart
	err := r.db.Where("product_id = ?", productID).Find(&carts).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []Cart{}, errors.New("Cart not found")
	}
	return carts, err
}

func (r *repository) FindStatusCardByUser(userID int, isActived string) ([]Cart, error) {
	var carts []Cart
	err := r.db.Where("user_id = ? AND isActived = ?", userID, isActived).Find(&carts).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []Cart{}, errors.New("Cart not found")
	}
	return carts, err
}

func (r *repository) SumTotalPriceByUser(userID int, isActived string) (int, error) {
	var total int
	err := r.db.Model(&Cart{}).
		Where("user_id = ? AND isActived = ?", userID, isActived).
		Select("SUM(total_price)").
		Scan(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCart(cart Cart) (Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) UpdateCart(cart Cart) (Cart, error) {
	err := r.db.Save(&cart).Error
	return cart, err
}

func (r *repository) DeleteCart(cart Cart) (Cart, error) {
	err := r.db.Delete(&cart).Error
	return cart, err
}
