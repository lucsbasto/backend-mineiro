package repositories

import (
	"github.com/lucsbasto/backend-mineiro/models"
	"gorm.io/gorm"
)

type SalesRepository interface {
	Create(sale *models.Sales) error
	FindById(saleId string) (*models.Sales, error)
	FindAll()([]models.Sales, error)
	FindMostRecent() (*models.Sales, error)
	Update(sale *models.Sales) error
	Delete(saleId string) error
}

type salesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) *salesRepository{
	return &salesRepository{db:db}
}

func (r *salesRepository) Create(sale *models.Sales) error {
	return r.db.Create(sale).Error
}

func (r *salesRepository) FindByID(saleId string) (*models.Sales, error){
	var sale models.Sales
	err := r.db.First(&sale, "id = ?", saleId).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *salesRepository) FindAll()([]models.Sales, error){
	var sales []models.Sales
	err := r.db.Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *salesRepository) FindMostRecent() (*models.Sales, error){
	var sale models.Sales
	err := r.db.Order("created_at desc").First(&sale).Error
	if err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *salesRepository) Update(sale *models.Sales) error {
	return r.db.Save(sale).Error
}

func (r *salesRepository) Delete(salesId string) error {
	return r.db.Unscoped().Delete(&models.Sales{}, "id = ?", salesId).Error
}
