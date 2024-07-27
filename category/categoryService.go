package category

import (
	"errors"

	"gorm.io/gorm"
)

type CategoryService interface {
	FindAll() ([]Category, error)
	FindCategoryByID(ID int) (Category, error)
	CreateCategory(category CategoryCreateRequest) (Category, error)
	UpdateCategory(ID int, category CategoryUpdateRequest) (Category, error)
	DeleteCategory(ID int) (Category, error)
}

type service struct {
	categoryRepository CategoryRepository
}

func NewService(categoryRepository CategoryRepository) *service {
	return &service{categoryRepository}
}

func (s *service) FindAll() ([]Category, error) {
	return s.categoryRepository.FindAll()
}

func (s *service) FindCategoryByID(ID int) (Category, error) {
	return s.categoryRepository.FindCategoryByID(ID)
}

func (s *service) CreateCategory(categoryRequest CategoryCreateRequest) (Category, error) {
	categoryData := Category{
		Name: categoryRequest.Name,
	}

	category, err := s.categoryRepository.CreateCategory(categoryData)

	return category, err
}

func (s *service) UpdateCategory(ID int, categoryRequest CategoryUpdateRequest) (Category, error) {
	category, err := s.categoryRepository.FindCategoryByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Category{}, errors.New("Category not found")
	}

	if err != nil {
		return Category{}, err
	}

	if categoryRequest.Name != "" {
		category.Name = categoryRequest.Name
	}

	return s.categoryRepository.UpdateCategory(category)
}

func (s *service) DeleteCategory(ID int) (Category, error) {
	category, err := s.categoryRepository.FindCategoryByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Category{}, errors.New("Category not found")
	}

	if err != nil {
		return Category{}, err
	}

	return s.categoryRepository.DeleteCategory(category)
}
