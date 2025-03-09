package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesProduct struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	SaleID    string    `gorm:"primaryKey;not null" json:"sale_id"`    // Chave estrangeira para Sales
	ProductID string    `gorm:"primaryKey;not null" json:"product_id"` // Chave estrangeira para Product
	Quantity  int       `json:"quantity"`  
	Sold      int       `json:"sold"`      
	Returned  int       `json:"returned"`  
	UnitCost  float64   `json:"unit_cost"` 
	Price     float64   `json:"price"`     
	TotalCost float64   `json:"total_cost"` 
	Revenue   float64   `json:"revenue"`   
	Profit    float64   `json:"profit"`    
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Sale      Sales     `gorm:"foreignKey:SaleID"`
	Product   Product   `gorm:"foreignKey:ProductID"`
}

func (sp *SalesProduct) BeforeUpdate(tx *gorm.DB) (err error) {
	sp.TotalCost = float64(sp.Sold) * sp.UnitCost
	sp.Revenue = float64(sp.Sold) * sp.Price
	sp.Profit = sp.Revenue - sp.TotalCost
	sp.UpdatedAt = time.Now()
	return nil
}

func (sp *SalesProduct) BeforeCreate(tx *gorm.DB) (err error) {
	sp.ID = uuid.New().String()
	sp.TotalCost = float64(sp.Sold) * sp.UnitCost
	sp.Revenue = float64(sp.Sold) * sp.Price
	sp.Profit = sp.Revenue - sp.TotalCost
	return nil
}
