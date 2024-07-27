package product

import (
	"errors"

	"gorm.io/gorm"
)

type ProductService interface {
	FindAll() ([]Product, error)
	FindProductByID(ID int) (Product, error)
	GetProductByUserIDAndCategoryID(userID int, categoryID int) ([]Product, error)
	GetProductByUser(userID int) ([]Product, error)
	GetProductByCategory(categoryID int) ([]Product, error)
	CreateProduct(product ProductCreateRequest) (Product, error)
	UpdateProduct(ID int, product ProductUpdateRequest) (Product, error)
	DeleteProduct(ID int) (Product, error)
}

type service struct {
	productRepository ProductRepository
}

func NewService(productRepository ProductRepository) *service {
	return &service{productRepository}
}

func (s *service) FindAll() ([]Product, error) {
	return s.productRepository.FindAll()
}

func (s *service) FindProductByID(ID int) (Product, error) {
	return s.productRepository.FindProductByID(ID)
}

func (s *service) GetProductByUserIDAndCategoryID(userID int, categoryID int) ([]Product, error) {
	return s.productRepository.GetProductByUserIDAndCategoryID(userID, categoryID)
}

func (s *service) GetProductByUser(userID int) ([]Product, error) {
	return s.productRepository.GetProductByUser(userID)
}

func (s *service) GetProductByCategory(categoryID int) ([]Product, error) {
	return s.productRepository.GetProductByCategory(categoryID)
}

func (s *service) CreateProduct(productRequest ProductCreateRequest) (Product, error) {
	productData := Product{
		UserID:      productRequest.UserID,
		CategoryID:  productRequest.CategoryID,
		Name:        productRequest.Name,
		Image:       productRequest.Image.Filename,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
	}

	product, err := s.productRepository.CreateProduct(productData)

	return product, err
}

func (s *service) UpdateProduct(ID int, productRequest ProductUpdateRequest) (Product, error) {
	product, err := s.productRepository.FindProductByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Product{}, errors.New("Product not found")
	}

	if err != nil {
		return Product{}, err
	}

	if productRequest.CategoryID != 0 {
		product.CategoryID = productRequest.CategoryID
	}
	if productRequest.Name != "" {
		product.Name = productRequest.Name
	}
	if productRequest.Image != nil {
		product.Image = productRequest.Image.Filename
	}
	if productRequest.Description != "" {
		product.Description = productRequest.Description
	}
	if productRequest.Price != 0 {
		product.Price = productRequest.Price
	}
	if productRequest.Stock != 0 {
		product.Stock = productRequest.Stock
	}

	return s.productRepository.UpdateProduct(product)
}

func (s *service) DeleteProduct(ID int) (Product, error) {
	product, err := s.productRepository.FindProductByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Product{}, errors.New("Product not found")
	}

	if err != nil {
		return Product{}, err
	}

	return s.productRepository.DeleteProduct(product)
}
