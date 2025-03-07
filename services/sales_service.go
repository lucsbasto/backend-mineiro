package services

import (
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

func (s *SalesService) CreateSale(sale *models.Sales) error {
	// Crie a venda
	if err := s.salesRepo.Create(sale); err != nil {
		return err
	}

	// Associe os produtos à venda
	for i := range sale.Products {
		// Atribua o ID da venda ao produto
		sale.Products[i].Sales = append(sale.Products[i].Sales, *sale)
	}

	// Atualize os produtos associados
	if err := s.salesRepo.UpdateProducts(sale.Products); err != nil {
		return err
	}

	return nil
}


func (s *SalesService) CreateSaleWithProducts(sale *models.Sales) error {
	tx, err := s.salesRepo.BeginTransaction()
	if err != nil {
			return err
	}
	defer tx.Rollback()

	// Crie a venda primeiro
	saleErr := s.salesRepo.CreateInTransaction(tx, sale)
	if saleErr != nil {
			return saleErr
	}

	// Agora que a venda foi criada, associamos os produtos
	for _, product := range sale.Products {
			// Crie o produto se ele não existir
			productErr := s.productRepo.CreateInTransaction(tx, &product)
			if productErr != nil {
					return productErr
			}
	}

	commitErr := tx.Commit().Error
	if commitErr != nil {
			return commitErr
	}

	return nil
}


func (s *SalesService) FindSaleByID(saleId string) (*models.Sales, error) {

	return s.salesRepo.FindByID(saleId)
}

func (s *SalesService) FindAll(isAdmin bool, userId string) ([]models.Sales, error) {
	return s.salesRepo.FindAll(isAdmin, userId)
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
