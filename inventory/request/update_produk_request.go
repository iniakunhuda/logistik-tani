package request

type UpdateProdukRequest struct {
	Description string `json:"description"`
	PriceBuy    int    `json:"price_buy"`
	PriceSell   int    `json:"price_sell"`
	Stock       int    `json:"stock"`
}
