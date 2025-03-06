package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Sales representa uma venda.
type Sales struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	TotalRevenue  float64   `json:"total_revenue"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	SalesProducts []SalesProduct `gorm:"foreignKey:SaleID" json:"sales_products"`
	UserID        string    `gorm:"type:varchar(255);not null" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID" json:"user"`
}

// BeforeCreate gera um ID único para a venda antes de criá-la.
func (sale *Sales) BeforeCreate(tx *gorm.DB) (err error) {
	sale.ID = uuid.New().String()
	return nil
}

