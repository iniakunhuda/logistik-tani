package request

import "time"

type CreatePayoutRequest struct {
	IDSender                int        `json:"id_sender" validate:"required"`
	IDReceiver              int        `json:"id_receiver" validate:"required"`
	TotalAmount             float64    `json:"total_amount" validate:"required"`
	BankNote                string     `json:"bank_note"`
	IDPurchaseReportsToBank int        `json:"id_purchase_reports_to_bank"`
	DatePayout              *time.Time `json:"date_payout"`
	Status                  string     `json:"status" validate:"required,oneof=pending approved rejected"`
	CreatedDate             time.Time  `json:"created_date"`
	RejectedDate            *time.Time `json:"rejected_date"`
	RejectedMessage         *string    `json:"rejected_message"`
	ApprovedMessage         *string    `json:"approved_message"`
	ApprovedDate            *time.Time `json:"approved_date"`
}
