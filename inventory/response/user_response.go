package response

import "time"

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	Telp     string `json:"telp" validate:"required"`
	Role     string `json:"role" validate:"required"`

	Saldo       uint   `json:"saldo"`
	LastLogin   string `json:"last_login"`
	AlamatKebun string `json:"alamat_kebun"`
	TotalObat   uint   `json:"total_obat"`
	TotalPupuk  uint   `json:"total_pupuk"`
	TotalBibit  uint   `json:"total_bibit"`
	TotalTebu   uint   `json:"total_tebu"`
	LuasLahan   uint   `json:"luas_lahan"`

	Token        *string    `json:"token"`
	TokenExpired *time.Time ` json:"token_expired"`

	CreatedAt time.Time ` json:"created_at"`
	UpdatedAt time.Time ` json:"updated_at"`
	Password  *struct{} `json:"password,omitempty"`
}
