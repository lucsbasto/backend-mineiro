package services

import (
	"github.com/lucsbasto/backend-mineiro/controllers/dtos"
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

func (s *SalesProductService) FindSalesByFormattedDate(date string, isAdmin bool, userId string) ([]dtos.SaleResponseDto, error) {
	salesProduct, error := s.repo.FindByFormattedDate(date, isAdmin, userId)
	if error != nil {
		return nil, error
	}
	var salesResponse []dtos.SaleResponseDto 
	for _, sale := range salesProduct {
		salesResponse = append(salesResponse, dtos.SaleResponseDto{
			ID:        sale.ID,
			SaleId:    sale.SaleID,
			Type: 	   sale.Product.Type,
			ProductId: sale.ProductID,
			Price:     sale.Price,
			Quantity:  sale.Quantity,
			Sold:      sale.Sold,
			Revenue:   sale.Revenue,
			Returned:  sale.Returned,
			UnitCost:  sale.UnitCost,
			
			TotalCost: CalculateTotalCost(&sale),
			Profit: 	CalculateProfit(&sale),
		})
	}
	return salesResponse, nil
}

func CalculateTotalCost(sale *models.SalesProduct) float64 {
	return sale.UnitCost * float64(sale.Quantity)
}

func CalculateProfit(sale *models.SalesProduct) float64 {
	totalPrice := sale.Price * float64(sale.Quantity)
	return totalPrice - (sale.UnitCost * float64(sale.Quantity))
}

func (s *SalesProductService) ListOne(id string) (*models.SalesProduct, error) {
	return s.repo.ListOne(id)
}

func (s *SalesProductService) Update(salesProduct *models.SalesProduct) error {
	return s.repo.Update(salesProduct)
}
