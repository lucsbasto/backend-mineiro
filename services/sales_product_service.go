package services

import (
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/repositories"
)

type SalesProductService struct {
	repo repositories.SalesProductRepository
}

func NewSalesProductService(repo repositories.SalesProductRepository) *SalesProductService {
	return &SalesProductService{repo: repo}
}

func (s *SalesProductService) CreateSalesProduct(salesProduct *models.SalesProduct) error {
	return s.repo.Create(salesProduct)
}
