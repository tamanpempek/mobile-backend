package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserService interface {
	FindAll() ([]User, error)
	FindUsersByRole(role string) ([]User, error)
	FindUserByID(ID any) (User, error)
	FindUserByEmail(email string) (User, error)
	CreateUser(user UserCreateRequest) (User, error)
	UpdateUser(ID int, user UserUpdateRequest) (User, error)
	DeleteUser(ID int) (User, error)
}

type service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *service {
	return &service{userRepository}
}

func (s *service) FindAll() ([]User, error) {
	return s.userRepository.FindAll()
}

func (s *service) FindUsersByRole(role string) ([]User, error) {
	return s.userRepository.FindUsersByRole(role)
}

func (s *service) FindUserByID(ID any) (User, error) {
	return s.userRepository.FindUserByID(ID)
}

func (s *service) FindUserByEmail(email string) (User, error) {
	return s.userRepository.FindUserByEmail(email)
}

func (s *service) CreateUser(userRequest UserCreateRequest) (User, error) {
	userData := User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Whatsapp: userRequest.Whatsapp,
		Gender:   userRequest.Gender,
		Role:     userRequest.Role,
	}

	user, err := s.userRepository.CreateUser(userData)

	return user, err
}

func (s *service) UpdateUser(ID int, userRequest UserUpdateRequest) (User, error) {
	user, err := s.userRepository.FindUserByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("User not found")
	}

	if err != nil {
		return User{}, err
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	if userRequest.Email != "" {
		user.Email = userRequest.Email
	}
	if userRequest.Password != "" {
		user.Password = userRequest.Password
	}
	if userRequest.Whatsapp != "" {
		user.Whatsapp = userRequest.Whatsapp
	}
	if userRequest.Gender != "" {
		user.Gender = userRequest.Gender
	}
	if userRequest.Role != "" {
		user.Role = userRequest.Role
	}

	return s.userRepository.UpdateUser(user)
}

func (s *service) DeleteUser(ID int) (User, error) {
	user, err := s.userRepository.FindUserByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("User not found")
	}

	if err != nil {
		return User{}, err
	}

	return s.userRepository.DeleteUser(user)
}
