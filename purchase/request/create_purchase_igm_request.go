package request

import "time"

type CreatePurchaseIgmRequest struct {
	PurchaseDate time.Time `json:"purchase_date" validate:"required"`
	Note         string    `json:"note"`
	TotalTebu    int       `json:"total_tebu" validate:"required"`
	TotalPrice   int       `json:"total_price" validate:"required"`
	TotalFarmer  int       `json:"total_farmer" validate:"required"`
	Status       string    `json:"status" validate:"required"`

	Items []CreatePurchaseIgmItemRequest `json:"items" validate:"required"`
}

type CreatePurchaseIgmItemRequest struct {
	IDUser       int `json:"id_user" validate:"required"`
	IDUserLand   int `json:"id_user_land" validate:"required"`
	IDProduction int `json:"id_production" validate:"required"`
	TotalKg      int `json:"total_kg" validate:"required"`
	HargaKg      int `json:"harga_kg" validate:"required"`
	Subtotal     int `json:"subtotal" validate:"required"`
}
