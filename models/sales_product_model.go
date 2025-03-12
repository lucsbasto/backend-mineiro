package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SalesProduct struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	SaleID    string    `gorm:"primaryKey;not null" json:"saleId"`    // Chave estrangeira para Sales
	ProductID string    `gorm:"primaryKey;not null" json:"productId"` // Chave estrangeira para Product
	Quantity  int       `json:"quantity"`  
	Sold      int       `json:"sold"`      
	Returned  int       `json:"returned"`  
	UnitCost  float64   `json:"unitCost"` 
	Price     float64   `json:"price"`     
	TotalCost float64   `json:"totalCost"` 
	Revenue   float64   `json:"revenue"`   
	Profit    float64   `json:"profit"`    
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
