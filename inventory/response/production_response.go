package response

import (
	"github.com/iniakunhuda/logistik-tani/inventory/model"
)

type ProductionResponse struct {
	model.Production
	UserDetail UserResponse        `json:"user_detail"`
	LandDetail UserLandRowResponse `json:"land_detail"`
}

type UserLandRowResponse struct {
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
