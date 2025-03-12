package services

import (
	"github.com/lucsbasto/backend-mineiro/controllers/dtos"
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/repositories"
)

type SalesService struct {
	salesRepo       repositories.SalesRepository
	productRepo repositories.ProductRepository
}

func NewSalesService(salesRepo repositories.SalesRepository, productRepo repositories.ProductRepository) *SalesService {
	return &SalesService{salesRepo: salesRepo, productRepo: productRepo}
}

func (s *SalesService) CreateSale(dto *dtos.CreateSaleDTO, userID string) error {
	// Criar a venda
	sale := models.Sales{
		UserID:    userID,  // Agora associamos o userID à venda
		TotalRevenue: 0.0,  // Total inicial, pode ser calculado depois
	}
	if err := s.salesRepo.Create(&sale); err != nil {
		return err
	}

	for _, productDTO := range dto.Products {
		product := models.Product{
			ID:        productDTO.ProductID,
		}

		// Cria o produto e associa à venda
		salesProduct := models.SalesProduct{
			SaleID:    sale.ID,
			ProductID: product.ID,
			Quantity:  productDTO.Quantity,
			UnitCost:  productDTO.UnitCost,
			Price:     productDTO.Price,
		}

		if err := s.salesRepo.CreateSalesProduct(&salesProduct); err != nil {
			return err
		}
	}
	return nil
}

func (s *SalesService) FindSaleByID(saleId string) (*models.Sales, error) {

	return s.salesRepo.FindByID(saleId)
}

func (s *SalesService) FindAll(isAdmin bool, userId string) ([]models.Sales, error) {
	return s.salesRepo.FindAll(isAdmin, userId)
}

func (s *SalesService) UpdateSale(sale *models.Sales) error {
	return s.salesRepo.Update(sale)
}

// func (s *SalesService) DeleteSale(saleId string) error {
// 	return s.salesRepo.Delete(saleId)
// }
