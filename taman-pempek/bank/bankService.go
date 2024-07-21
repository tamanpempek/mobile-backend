package bank

import (
	"errors"

	"gorm.io/gorm"
)

type BankService interface {
	FindAll() ([]Bank, error)
	FindAdminBanks() ([]Bank, error)
	FindBanksByUser(userID int) ([]Bank, error)
	FindBankByID(ID int) (Bank, error)
	CreateBank(bank BankCreateRequest) (Bank, error)
	UpdateBank(ID int, bank BankUpdateRequest) (Bank, error)
	DeleteBank(ID int) (Bank, error)
}

type service struct {
	bankRepository BankRepository
}

func NewService(bankRepository BankRepository) *service {
	return &service{bankRepository}
}

func (s *service) FindAll() ([]Bank, error) {
	return s.bankRepository.FindAll()
}

func (s *service) FindAdminBanks() ([]Bank, error) {
	return s.bankRepository.FindAdminBanks()
}

func (s *service) FindBanksByUser(userID int) ([]Bank, error) {
	return s.bankRepository.FindBanksByUser(userID)
}

func (s *service) FindBankByID(ID int) (Bank, error) {
	return s.bankRepository.FindBankByID(ID)
}

func (s *service) CreateBank(bankRequest BankCreateRequest) (Bank, error) {
	bankData := Bank{
		UserID: bankRequest.UserID,
		Type:   bankRequest.Type,
		Name:   bankRequest.Name,
		Number: bankRequest.Number,
	}

	bank, err := s.bankRepository.CreateBank(bankData)

	return bank, err
}

func (s *service) UpdateBank(ID int, bankRequest BankUpdateRequest) (Bank, error) {
	bank, err := s.bankRepository.FindBankByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Bank{}, errors.New("Bank not found")
	}

	if err != nil {
		return Bank{}, err
	}

	if bankRequest.Type != "" {
		bank.Type = bankRequest.Type
	}
	if bankRequest.Name != "" {
		bank.Name = bankRequest.Name
	}
	if bankRequest.Number != "" {
		bank.Number = bankRequest.Number
	}

	return s.bankRepository.UpdateBank(bank)
}

func (s *service) DeleteBank(ID int) (Bank, error) {
	bank, err := s.bankRepository.FindBankByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Bank{}, errors.New("Bank not found")
	}

	if err != nil {
		return Bank{}, err
	}

	return s.bankRepository.DeleteBank(bank)
}
