package panenresponse

import (
	"time"
	// "github.com/iniakunhuda/logistik-tani/purchase/response"
	// userresponse "github.com/iniakunhuda/logistik-tani/purchase/response/user_response"
)

type ProductionListResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    []ProductionRowResponse `json:"data"`
}

type ProductionDetailResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    ProductionRowResponse `json:"data"`
}

type ProductionRowResponse struct {
	ID           uint       `json:"id"`
	IDUser       int        `json:"id_user"`
	IDUserLand   int        `json:"id_user_land"`
	Title        string     `json:"title"`
	DateMonth    int        `json:"date_month"`
	DateYear     int        `json:"date_year"`
	DateStart    time.Time  `json:"date_start"`
	DateEnd      *time.Time `json:"date_end"`
	TotalPanenKg int        `json:"total_panen_kg"`
	Status       string     `json:"status"`

	Histories []ProductionHistoryResponse `json:"histories"`
	// UserDetail ProductionUserResponse        `json:"user_detail"`
	// LandDetail ProductionUserLandRowResponse `json:"land_detail"`
}

type ProductionHistoryResponse struct {
	ID             uint                           `json:"id"`
	IDProduction   int                            `json:"id_production"`
	IDProductOwner int                            `json:"id_product_owner"`
	QtyUse         int                            `json:"qty_use"`
	Note           string                         `json:"note"`
	Date           time.Time                      `json:"date"`
	ProductOwner   ProductionProductOwnerResponse `json:"product_owner"`
}

type ProductionProductOwnerResponse struct {
	ID          uint   `json:"id"`
	IDUser      int    `json:"id_user"`
	IDProduct   int    `json:"id_product"`
	Stock       int    `json:"stock"`
	PriceBuy    int    `json:"price_buy"`
	PriceSell   int    `json:"price_sell"`
	Description string `json:"description"`

	Product ProductionProductResponse `json:"product"`
}

type ProductionProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	PriceBuy    int    `json:"price_buy"`
	PriceSell   int    `json:"price_sell"`
}

type ProductionUserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address" validate:"required"`
	Telp    string `json:"telp" validate:"required"`
	Role    string `json:"role" validate:"required"`
	Saldo   uint   `json:"saldo"`
}

type ProductionUserLandRowResponse struct {
	ID          uint    `json:"id"`
	IDUser      uint    `json:"id_user"`
	LandName    string  `json:"land_name"`
	LandAddress string  `json:"land_address"`
	LandArea    float64 `json:"land_area"`
	TotalObat   float64 `json:"total_obat"`
	TotalPupuk  float64 `json:"total_pupuk"`
	TotalBibit  float64 `json:"total_bibit"`
	TotalTebu   float64 `json:"total_tebu"`
}
