package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product representa um produto.
type Product struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	SalesProducts []SalesProduct `gorm:"foreignKey:ProductID" json:"sales_products"`
}

// BeforeCreate gera um ID único para o produto antes de criá-lo.
func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New().String()
	return nil
}
