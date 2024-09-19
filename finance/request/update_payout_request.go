package request

import "time"

type UpdatePayoutRequest struct {
	Status  string `json:"status" validate:"required"`
	Message string `json:"message"`

	TotalAmount *float64   `json:"total_amount"`
	BankNote    *string    `json:"bank_note"`
	DatePayout  *time.Time `json:"date_payout"`
}
