package services

import (
	"github.com/lucsbasto/backend-mineiro/models"
	"github.com/lucsbasto/backend-mineiro/repositories"
)

type SalesService struct {
	salesRepo       repositories.SalesRepository
	salesProductRepo repositories.SalesProductRepository
}

func NewSalesService(salesRepo repositories.SalesRepository, salesProductRepo repositories.SalesProductRepository) *SalesService {
	return &SalesService{salesRepo: salesRepo, salesProductRepo: salesProductRepo}
}

func (s *SalesService) CreateSale(sale *models.Sales) error {
	if err := s.salesRepo.Create(sale); err != nil {
		return err
	}

	for _, sp := range sale.SalesProducts {
		sp.SaleID = sale.ID
		if err := s.salesProductRepo.Create(&sp); err != nil {
			return err
		}
	}

	return nil
}

func (s *SalesService) CreateSaleWithProducts(sale *models.Sales) error {
	tx, err := s.salesRepo.BeginTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	saleErr := s.salesRepo.CreateInTransaction(tx, sale)
	for _, sp := range sale.SalesProducts {
		sp.SaleID = sale.ID
		spErr := s.salesProductRepo.CreateInTransaction(tx, &sp)
		if spErr != nil {
			return spErr
		}
	}

	commitErr := tx.Commit().Error

	if saleErr != nil {
		return saleErr
	}

	if commitErr != nil {
		return commitErr
	}

	return nil
}

func (s *SalesService) FindSaleByID(saleId string) (*models.Sales, error) {

	return s.salesRepo.FindByID(saleId)
}

func (s *SalesService) FindAll() ([]models.Sales, error) {
	return s.salesRepo.FindAll()
}

func (s *SalesService) FindSalesByFormattedDate(date string) ([]models.Sales, error) {
	return s.salesRepo.FindByFormattedDate(date)
}

func (s *SalesService) UpdateSale(sale *models.Sales) error {
	return s.salesRepo.Update(sale)
}

func (s *SalesService) DeleteSale(saleId string) error {
	return s.salesRepo.Delete(saleId)
}
