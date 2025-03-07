package repositories

import (
	"github.com/lucsbasto/backend-mineiro/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(productId string) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(productId string) error
	CreateInTransaction(tx *gorm.DB, product *models.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindByID(productId string) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, "id = ?", productId).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(productId string) error {
	return r.db.Unscoped().Delete(&models.Product{}, "id = ?", productId).Error
}

func (r *productRepository) CreateInTransaction(tx *gorm.DB, sale *models.Product) error {
	return tx.Create(sale).Error
}
