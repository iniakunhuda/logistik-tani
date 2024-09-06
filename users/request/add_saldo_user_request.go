package request

type AddSaldoUserRequest struct {
	NewSaldo int `json:"new_saldo" validate:"required"`
}
