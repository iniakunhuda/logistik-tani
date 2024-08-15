package request

type UpdateSalesRequest struct {
	IDPenjual uint   `json:"id_penjual" validate:"required"` // ID pembibit
	Status    string `json:"status" validate:"required"`
}
