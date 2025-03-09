package types

type UpdateSalesProduct struct {
	ID       string   `json:"id,omitempty"`
	Quantity *int     `json:"quantity,omitempty"`
	Sold     *int     `json:"sold,omitempty"`
	Returned *int     `json:"returned,omitempty"`
	UnitCost *float64 `json:"unitCost,omitempty"`
	Price    *float64 `json:"price,omitempty"`
}
