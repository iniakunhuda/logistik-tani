package request

type CreateUserRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=20"`
	Alamat   string `json:"alamat" validate:"required"`
	Telp     string `json:"telp" validate:"required"`
	Role     string `json:"role" validate:"required"`

	Saldo       uint   `json:"saldo"`
	// LastLogin   string `json:"last_login"`
	AlamatKebun string `json:"alamat_kebun"`
	TotalObat   uint   `json:"total_obat"`
	TotalPupuk  uint   `json:"total_pupuk"`
	TotalBibit  uint   `json:"total_bibit"`
	TotalTebu   uint   `json:"total_tebu"`
	LuasLahan   uint   `json:"luas_lahan"`
}
