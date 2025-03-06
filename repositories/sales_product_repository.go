package repositories

import (
	"gorm.io/gorm"

	"github.com/lucsbasto/backend-mineiro/models"
)

type SalesProductRepository interface {
	Create(salesProduct *models.SalesProduct) error
	BeginTransaction() (*gorm.DB, error)
	CreateInTransaction(tx *gorm.DB, salesProduct *models.SalesProduct) error
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}

type salesProductRepository struct {
	db *gorm.DB
}

func NewSalesProductRepository(db *gorm.DB) *salesProductRepository {
	return &salesProductRepository{db: db}
}

func (r *salesProductRepository) Create(salesProduct *models.SalesProduct) error {
	return r.db.Create(salesProduct).Error
}

func (r *salesProductRepository) BeginTransaction() (*gorm.DB, error) {
	return r.db.Begin(), nil
}

func (r *salesProductRepository) CreateInTransaction(tx *gorm.DB, salesProduct *models.SalesProduct) error {
	return tx.Create(salesProduct).Error
}

func (r *salesProductRepository) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *salesProductRepository) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}
