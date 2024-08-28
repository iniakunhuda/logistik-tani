package request

type UpdateSalesRequest struct {
	IDSeller uint   `json:"id_seller" validate:"required"`
	Status   string `json:"status" validate:"required"`
}
