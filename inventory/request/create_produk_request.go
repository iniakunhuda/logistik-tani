package request

type CreateProdukRequest struct {
	ID          uint   `json:"id"`
	IDUser      uint   `json:"id_user" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	PriceBuy    int    `json:"price_buy" validate:"required"`
	PriceSell   int    `json:"price_sell" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
}
