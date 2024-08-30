package model

import (
	"time"
)

type PayoutHistory struct {
	ID                      uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	IDSender                int        `gorm:"not null" json:"id_sender" validate:"required"`
	IDReceiver              int        `gorm:"not null" json:"id_receiver" validate:"required"`
	NoInvoice               string     `gorm:"type:varchar(255);not null" json:"no_invoice" validate:"required"`
	TotalAmount             float64    `gorm:"type:decimal(10,2);not null" json:"total_amount" validate:"required"`
	BankNote                string     `gorm:"type:text" json:"bank_note"`
	IDPurchaseReportsToBank int        `json:"id_purchase_reports_to_bank"`
	DatePayout              *time.Time `gorm:"type:date" json:"date_payout"`
	Status                  string     `gorm:"type:enum('pending', 'approved', 'rejected');not null" json:"status" validate:"required,oneof=pending approved rejected"`
	CreatedDate             time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_date"`
	RejectedDate            *time.Time `gorm:"type:timestamp" json:"rejected_date"`
	RejectedMessage         *string    `gorm:"type:text" json:"rejected_message"`
	ApprovedMessage         *string    `gorm:"type:text" json:"approved_message"`
	ApprovedDate            *time.Time `gorm:"type:timestamp" json:"approved_date"`
}

func (PayoutHistory) TableName() string {
	return "payout_history"
}
