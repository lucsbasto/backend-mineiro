package models

import (
	"time"

	"gorm.io/gorm"
)

// SalesProduct representa uma linha de venda de um produto.
type SalesProduct struct {
	SaleID    string    `gorm:"primaryKey;not null" json:"sale_id"`
	ProductID string    `gorm:"primaryKey;not null" json:"product_id"`
	Quantity  int       `json:"quantity"`  // Quantidade do produto na venda
	Sold      int       `json:"sold"`      // Quantidade vendida
	Returned  int       `json:"returned"`  // Quantidade retornada
	UnitCost  float64   `json:"unit_cost"` // Custo unitário do produto
	Price     float64   `json:"price"`     // Preço do produto nesta venda
	Total     float64   `json:"total"`     // Total da linha de venda (Sold * Price)
	TotalCost float64   `json:"total_cost"`// Total do custo (Sold * UnitCost)
	Revenue   float64   `json:"revenue"`   // Receita (Total)
	Profit    float64   `json:"profit"`    // Lucro (Revenue - TotalCost)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Sale      Sales     `gorm:"foreignKey:SaleID" json:"-"`
	Product   Product   `gorm:"foreignKey:ProductID"`
}

// BeforeCreate calcula os totais baseados nas quantidades de produtos vendidos.
func (sp *SalesProduct) BeforeCreate(tx *gorm.DB) (err error) {
	sp.Total = float64(sp.Sold) * sp.Price
	sp.TotalCost = float64(sp.Sold) * sp.UnitCost
	sp.Revenue = sp.Total
	sp.Profit = sp.Revenue - sp.TotalCost
	return nil
}
