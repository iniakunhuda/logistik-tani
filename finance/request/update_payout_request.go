package request

type UpdatePayoutRequest struct {
	Status  string `json:"status" validate:"required"`
	Message string `json:"message"`
}
