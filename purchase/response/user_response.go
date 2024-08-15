package response

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Alamat   string `json:"alamat"`
	Telp     string `json:"telp"`
	Role     string `json:"role"`

	Saldo       uint   `json:"saldo"`
	// LastLogin   string `json:"last_login"`
	AlamatKebun string `json:"alamat_kebun"`
	TotalObat   uint   `json:"total_obat"`
	TotalPupuk  uint   `json:"total_pupuk"`
	TotalBibit  uint   `json:"total_bibit"`
	TotalTebu   uint   `json:"total_tebu"`
	LuasLahan   uint   `json:"luas_lahan"`
}
