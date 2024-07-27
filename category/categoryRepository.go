package category

import (
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]Category, error)
	FindCategoryByID(ID int) (Category, error)
	CreateCategory(category Category) (Category, error)
	UpdateCategory(category Category) (Category, error)
	DeleteCategory(category Category) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *repository) FindCategoryByID(ID int) (Category, error) {
	var category Category
	err := r.db.First(&category, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Category{}, errors.New("Category not found")
	}
	return category, err
}

func (r *repository) CreateCategory(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *repository) UpdateCategory(category Category) (Category, error) {
	err := r.db.Save(&category).Error
	return category, err
}

func (r *repository) DeleteCategory(category Category) (Category, error) {
	err := r.db.Delete(&category).Error
	return category, err
}
