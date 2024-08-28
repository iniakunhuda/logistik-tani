package request

type CreateUserRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=20"`
	Address  string `json:"address" validate:"required"`
	Telp     string `json:"telp" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Saldo    uint   `json:"saldo"`
}
