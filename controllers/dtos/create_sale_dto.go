package dtos

type CreateSaleDTO struct {
	Products []struct {
		ProductID string  `json:"product_id"`
		Quantity  int     `json:"quantity"`
		Sold      int     `json:"sold"`
		Returned  int     `json:"returned"`
		UnitCost  float64 `json:"unit_cost"`
		Price     float64 `json:"price"`
	} `json:"products"`
}
