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

func (s *SalesProductService) FindAll() ([]models.SalesProduct, error) {
	return s.repo.FindAll()
}	

func (s *SalesProductService) FindSalesByFormattedDate(date string, isAdmin bool, userId string) ([]models.SalesProduct, error) {
	return s.repo.FindByFormattedDate(date, isAdmin, userId)
}