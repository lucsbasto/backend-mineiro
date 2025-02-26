package models

import "time"

type Product struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Type      string    `json:"type"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	Sold      *int      `json:"sold"`
	Returned  *int      `json:"returned"`
	UnitCost  float64   `json:"unit_cost"`
	Revenue   *float64  `json:"revenue"`
	TotalCost *float64  `json:"total_cost"`
	Profit    *float64  `json:"profit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}