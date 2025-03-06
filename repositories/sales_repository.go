package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/lucsbasto/backend-mineiro/models"
)

type SalesRepository interface {
	Create(sale *models.Sales) error
	FindByID(saleId string) (*models.Sales, error)
	FindAll() ([]models.Sales, error)
	FindByFormattedDate(date string) ([]models.Sales, error)
	Update(sale *models.Sales) error
	Delete(saleId string) error
	GetSaleWithProducts(saleID string) (*models.Sales, error)
	BeginTransaction() (*gorm.DB, error)
	CreateInTransaction(tx *gorm.DB, sale *models.Sales) error
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}

type salesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) *salesRepository {
	return &salesRepository{db: db}
}

func (r *salesRepository) Create(sale *models.Sales) error {
	return r.db.Create(sale).Error
}

func (r *salesRepository) FindByID(saleId string) (*models.Sales, error) {
	var sale models.Sales
	if err := r.db.Preload("User").Preload("SalesProducts").Preload("SalesProducts.Product").First(&sale, "id = ?", saleId).Error; err != nil {
		return nil, err
	}
	fmt.Printf("Sale: %+v\n", sale)
	return &sale, nil
}

func (r *salesRepository) FindAll() ([]models.Sales, error) {
	var sales []models.Sales
	if err := r.db.Preload("User").Preload("SalesProducts").Preload("SalesProducts.Product").
		Find(&sales).Error; err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *salesRepository) FindByFormattedDate(date string) ([]models.Sales, error) {
	var sales []models.Sales
	err := r.db.Preload("User").Preload("SalesProducts").Preload("SalesProducts.Product").
		Where("TO_CHAR(created_at, 'YYYY-MM-DD') = ?", date).
		Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
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
