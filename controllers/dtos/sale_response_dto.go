package dtos

type SaleResponseDto struct {
	SaleId    string  `json:"id"`
	ProductId string  `json:"product_id"`
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Sold      int     `json:"sold"`
	Returned  int     `json:"returned"`
	UnitCost  float64 `json:"unit_cost"`
	TotalCost float64 `json:"total_cost"`
	Profit    float64 `json:"profit"`
}
