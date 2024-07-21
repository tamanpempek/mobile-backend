package bank

import (
	"errors"

	"gorm.io/gorm"
)

type BankRepository interface {
	FindAll() ([]Bank, error)
	FindAdminBanks() ([]Bank, error)
	FindBanksByUser(userID int) ([]Bank, error)
	FindBankByID(ID int) (Bank, error)
	CreateBank(bank Bank) (Bank, error)
	UpdateBank(bank Bank) (Bank, error)
	DeleteBank(bank Bank) (Bank, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Bank, error) {
	var banks []Bank
	err := r.db.Find(&banks).Error
	return banks, err
}

func (r *repository) FindAdminBanks() ([]Bank, error) {
	var banks []Bank
	userID := 10
	err := r.db.Where("user_id = ?", userID).Find(&banks).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []Bank{}, errors.New("Bank not found")
	}
	return banks, err
}

func (r *repository) FindBanksByUser(userID int) ([]Bank, error) {
	var banks []Bank
	err := r.db.Where("user_id = ?", userID).Find(&banks).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []Bank{}, errors.New("Bank not found")
	}
	return banks, err
}

func (r *repository) FindBankByID(ID int) (Bank, error) {
	var bank Bank
	err := r.db.First(&bank, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Bank{}, errors.New("Bank not found")
	}
	return bank, err
}

func (r *repository) CreateBank(bank Bank) (Bank, error) {
	err := r.db.Create(&bank).Error
	return bank, err
}

func (r *repository) UpdateBank(bank Bank) (Bank, error) {
	err := r.db.Save(&bank).Error
	return bank, err
}

func (r *repository) DeleteBank(bank Bank) (Bank, error) {
	err := r.db.Delete(&bank).Error
	return bank, err
}
