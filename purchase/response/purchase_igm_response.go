package response

import "time"

type PurchaseIgmResponse struct {
	ID           uint      `json:"id"`
	NoInvoice    string    `json:"no_invoice"`
	PurchaseDate time.Time `json:"purchase_date"`
	Note         string    `json:"note"`
	TotalTebu    float64   `json:"total_tebu"`
	TotalPrice   float64   `json:"total_price"`
	TotalFarmer  int       `json:"total_farmer"`
	Status       string    `json:"status" validate:"required,oneof=open pending done"`

	Items []PurchaseIgmItemResponse `json:"items"`
}

type PurchaseIgmItemResponse struct {
	IDUser         int                           `json:"id_user"`
	IDUserLand     int                           `json:"id_user_land"`
	IDProduction   int                           `json:"id_production"`
	TotalKg        float64                       `json:"total_kg"`
	HargaKg        float64                       `json:"harga_kg"`
	Subtotal       float64                       `json:"subtotal"`
	UserDetail     PurchaseIgmUserDetailResponse `json:"user_detail"`
	UserLandDetail PurchaseIgmUserLandResponse   `json:"user_land_detail"`
}

type PurchaseIgmUserDetailResponse struct {
	UserResponse
}

type PurchaseIgmUserLandResponse struct {
	LandName    string  `json:"land_name"`
	LandAddress string  `json:"land_address"`
	LandArea    float64 `json:"land_area"`
	TotalObat   float64 `json:"total_obat"`
	TotalPupuk  float64 `json:"total_pupuk"`
	TotalBibit  float64 `json:"total_bibit"`
	TotalTebu   float64 `json:"total_tebu"`
}
