package response

type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Telp   string `json:"telp" validate:"required"`
	Role   string `json:"role" validate:"required"`
	Saldo  uint   `json:"saldo"`
}
