package services

import (
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/repositories"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) FindProductByID(productId string) (*models.Product, error) {
	return s.repo.FindByID(productId)
}

func (s *ProductService) FindAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(productId string) error {
	return s.repo.Delete(productId)
}
