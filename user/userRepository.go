package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]User, error)
	FindUsersByRole(role string) ([]User, error)
	FindUserByID(ID any) (User, error)
	FindUserByEmail(email string) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindUsersByRole(role string) ([]User, error) {
	var users []User
	err := r.db.Where("role = ?", role).Find(&users).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []User{}, errors.New("User not found")
	}
	return users, err
}

func (r *repository) FindUserByID(ID any) (User, error) {
	var user User
	err := r.db.First(&user, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("User not found")
	}
	return user, err
}

func (r *repository) FindUserByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("User not found")
	}
	return user, err
}

func (r *repository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) DeleteUser(user User) (User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
