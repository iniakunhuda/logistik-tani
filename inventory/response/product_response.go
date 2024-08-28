package response

type ProductResponse struct {
	ID          uint   `json:"id"`
	IDUser      uint   `json:"id_user"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PriceBuy    uint   `json:"price_buy"`
	PriceSell   uint   `json:"price_sell"`
	Category    string `json:"category"`
	Stock       uint   `json:"stock"`
}
