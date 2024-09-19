package request

type UpdateUserRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Telp    string `json:"telp" validate:"required"`
	Role    string `json:"role" validate:"required"`
	Address string `json:"address" validate:"required"`

	Saldo        uint    `json:"saldo"`
	Token        *string `json:"token"`
	TokenExpired *string `json:"token_expired"`
}
