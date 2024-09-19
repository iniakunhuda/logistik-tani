package request

type UpdateUserLandRequest struct {
	IDUser      uint    `json:"id_user" validate:"required"`
	LandName    string  `json:"land_name" validate:"required"`
	LandAddress string  `json:"land_address" validate:"required"`
	LandArea    float64 `json:"land_area" validate:"required"`
	TotalObat   float64 `json:"total_obat" validate:"required"`
	TotalPupuk  float64 `json:"total_pupuk" validate:"required"`
	TotalBibit  float64 `json:"total_bibit" validate:"required"`
	TotalTebu   float64 `json:"total_tebu" validate:"required"`
}
