package repositories

import (
	"gorm.io/gorm"

	"github.com/lucsbasto/backend-mineiro/models"
)

type SalesRepository interface {
	Create(sale *models.Sales) error
	FindByID(saleId string) (*models.Sales, error)
	FindAll(isAdmin bool, userId string) ([]models.Sales, error)
	Update(sale *models.Sales) error
	Delete(saleId string) error
	GetSaleWithProducts(saleID string) (*models.Sales, error)
	BeginTransaction() (*gorm.DB, error)
	CreateInTransaction(tx *gorm.DB, sale *models.Sales) error
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
	UpdateProducts(products []models.Product) error
	CreateSalesProduct(salesProduct *models.SalesProduct) error
}

type salesRepository struct {
	db *gorm.DB
}

func (r *salesRepository) UpdateProducts(products []models.Product) error {
	for _, product := range products {
		if err := r.db.Model(&product).Association("Sales").Append(product.Sales); err != nil {
			return err
		}
	}
	return nil
}


func NewSalesRepository(db *gorm.DB) *salesRepository {
	return &salesRepository{db: db}
}

func (r *salesRepository) Create(sale *models.Sales) error {
	return r.db.Create(sale).Error
}
// FindAll busca todas as vendas, com o relacionamento Many-to-Many para produtos
func (r *salesRepository) FindAll(isAdmin bool, userId string) ([]models.Sales, error) {
	var sales []models.Sales

	if isAdmin {
		// Admin pode ver todas as vendas
		if err := r.db.Preload("User").Preload("Products").Find(&sales).Error; err != nil {
			return nil, err
		}
	} else {
		// Usuário normal vê apenas suas vendas
		if err := r.db.Preload("User").Preload("Products").
			Where("user_id = ?", userId).
			Find(&sales).Error; err != nil {
			return nil, err
		}
	}
	return sales, nil
}

// FindByID busca uma venda específica com seus produtos associados
func (r *salesRepository) FindByID(saleId string) (*models.Sales, error) {
	var sale models.Sales
	if err := r.db.Preload("User").Preload("Products").
		First(&sale, "id = ?", saleId).Error; err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *salesRepository) Update(sale *models.Sales) error {
	return r.db.Save(sale).Error
}

func (r *salesRepository) Delete(saleId string) error {
	return r.db.Unscoped().Delete(&models.Sales{}, "id = ?", saleId).Error
}

func (r *salesRepository) GetSaleWithProducts(saleID string) (*models.Sales, error) {
	var sale models.Sales
	result := r.db.Preload("SalesProducts").Preload("SalesProducts.Product").First(&sale, "id = ?", saleID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}

func (r *salesRepository) BeginTransaction() (*gorm.DB, error) {
	return r.db.Begin(), nil
}

func (r *salesRepository) CreateInTransaction(tx *gorm.DB, sale *models.Sales) error {
	return tx.Create(sale).Error
}

func (r *salesRepository) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *salesRepository) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *salesRepository) CreateSalesProduct(salesProduct *models.SalesProduct) error {
	return r.db.Create(salesProduct).Error
}
