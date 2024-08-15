package request

type UpdatePurchaseRequest struct {
	IDPembeli uint   `json:"id_pembeli" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
