package dtos

type SaleResponseDto struct {
	ID        string  `json:"id"`
	SaleId    string  `json:"saleId"`
	ProductId string  `json:"productId"`
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Sold      int     `json:"sold"`
	Returned  int     `json:"returned"`
	UnitCost  float64 `json:"unitCost"`
	Revenue   float64 `json:"revenue"`
	TotalCost float64 `json:"totalCost"`
	Profit    float64 `json:"profit"`
}
