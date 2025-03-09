package repositories

import (
	"gorm.io/gorm"

	"github.com/lucsbasto/backend-mineiro/models"
)

type SalesProductRepository interface {
	FindAll() ([]models.SalesProduct, error)
	FindByFormattedDate(date string, isAdmin bool, userId string) ([]models.SalesProduct, error)
	Update(sale *models.SalesProduct) error
	ListOne(id string) (*models.SalesProduct, error)
}

type salesProductRepository struct {
	db *gorm.DB
}

func NewSalesProductRepository(db *gorm.DB) *salesProductRepository {
	return &salesProductRepository{db: db}
}

func (r *salesProductRepository) FindAll() ([]models.SalesProduct, error) {
	var salesProducts []models.SalesProduct
	if err := r.db.Preload("Product").Preload("Sale").Find(&salesProducts).Error; err != nil {
		return nil, err
	}
	return salesProducts, nil
}

func (r *salesProductRepository) FindByFormattedDate(date string, isAdmin bool, userId string) ([]models.SalesProduct, error) {
	var salesProduct []models.SalesProduct
	query := r.db.Preload("Product").Preload("Sale").Preload("Sale.User")

	if !isAdmin {
		query = query.Joins("JOIN sales ON sales.id = sales_products.sale_id").
			Where("sales.user_id = ?", userId)
	}

	err := query.Where("TO_CHAR(sales_products.created_at, 'YYYY-MM-DD') = ?", date).
		Find(&salesProduct).Error

	if err != nil {
		return nil, err
	}
	return salesProduct, nil
}

func (r *salesProductRepository) Update(sale *models.SalesProduct) error {
	return r.db.Model(sale).Omit("Sale", "Product").Updates(sale).Error
}

func (r *salesProductRepository) ListOne(id string) (*models.SalesProduct, error) {
	var sale models.SalesProduct
	if err := r.db.Preload("Product").Preload("Sale").First(&sale, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &sale, nil
}