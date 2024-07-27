package cart

import (
	"errors"

	"gorm.io/gorm"
)

type CartService interface {
	FindAll() ([]Cart, error)
	FindCartByID(ID int) (Cart, error)
	FindCartsByPaymentID(paymentID int) ([]Cart, error)
	FindCartsByProductID(productID int) ([]Cart, error)
	FindStatusCardByUser(userID int, isActived string) ([]Cart, error)
	SumTotalPriceByUser(userID int, isActived string) (int, error)
	CreateCart(cart CartCreateRequest) (Cart, error)
	UpdateCart(ID int, cart CartUpdateRequest) (Cart, error)
	DeleteCart(ID int) (Cart, error)
}

type service struct {
	cartRepository CartRepository
}

func NewService(cartRepository CartRepository) *service {
	return &service{cartRepository}
}

func (s *service) FindAll() ([]Cart, error) {
	return s.cartRepository.FindAll()
}

func (s *service) FindCartByID(ID int) (Cart, error) {
	return s.cartRepository.FindCartByID(ID)
}

func (s *service) FindCartsByPaymentID(paymentID int) ([]Cart, error) {
	return s.cartRepository.FindCartsByPaymentID(paymentID)
}

func (s *service) FindCartsByProductID(productID int) ([]Cart, error) {
	return s.cartRepository.FindCartsByProductID(productID)
}

func (s *service) FindStatusCardByUser(userID int, isActived string) ([]Cart, error) {
	return s.cartRepository.FindStatusCardByUser(userID, isActived)
}

func (s *service) SumTotalPriceByUser(userID int, isActived string) (int, error) {
	return s.cartRepository.SumTotalPriceByUser(userID, isActived)
}

func (s *service) CreateCart(cartRequest CartCreateRequest) (Cart, error) {
	cartData := Cart{
		UserID:     cartRequest.UserID,
		ProductID:  cartRequest.ProductID,
		PaymentID:  cartRequest.PaymentID,
		Quantity:   cartRequest.Quantity,
		TotalPrice: cartRequest.TotalPrice,
		IsActived:  cartRequest.IsActived,
	}

	cart, err := s.cartRepository.CreateCart(cartData)

	return cart, err
}

func (s *service) UpdateCart(ID int, cartRequest CartUpdateRequest) (Cart, error) {
	cart, err := s.cartRepository.FindCartByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Cart{}, errors.New("Cart not found")
	}

	if err != nil {
		return Cart{}, err
	}

	if cartRequest.ProductID != "" {
		cart.ProductID = cartRequest.ProductID
	}
	if cartRequest.PaymentID != "" {
		cart.PaymentID = cartRequest.PaymentID
	}
	if cartRequest.Quantity != "" {
		cart.Quantity = cartRequest.Quantity
	}
	if cartRequest.TotalPrice != "" {
		cart.TotalPrice = cartRequest.TotalPrice
	}
	if cartRequest.IsActived != "" {
		cart.IsActived = cartRequest.IsActived
	}

	return s.cartRepository.UpdateCart(cart)
}

func (s *service) DeleteCart(ID int) (Cart, error) {
	cart, err := s.cartRepository.FindCartByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Cart{}, errors.New("Cart not found")
	}

	if err != nil {
		return Cart{}, err
	}

	return s.cartRepository.DeleteCart(cart)
}
