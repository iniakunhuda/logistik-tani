package request

type UpdateProdukRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PriceBuy    int    `json:"price_buy"`
	PriceSell   int    `json:"price_sell"`
	Category    string `json:"category"`
}
