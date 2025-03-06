package repositories

import (
	"github.com/lucsbasto/backend-mineiro/models"
	"gorm.io/gorm"
)

// ProductRepository define as operações de repositório para produtos.
type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(productId string) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(productId string) error
}

// productRepository implementa ProductRepository.
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository cria uma nova instância de ProductRepository.
func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

// Create insere um novo produto no banco de dados.
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// FindByID busca um produto por ID.
func (r *productRepository) FindByID(productId string) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, "id = ?", productId).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// FindAll lista todos os produtos.
func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Update atualiza um produto existente.
func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete exclui um produto logicamente.
func (r *productRepository) Delete(productId string) error {
	return r.db.Unscoped().Delete(&models.Product{}, "id = ?", productId).Error
}
