package request

import "time"

type CreatePurchaseReportsToBankRequest struct {
	DateStart time.Time `json:"date_start" validate:"required"`
	DateEnd   time.Time `json:"date_end" validate:"required"`
	Note      string    `json:"note"`
	Status    string    `json:"status" validate:"required"`

	Purchases []CreatePurchaseReportsToBankItemRequest `json:"purchases" validate:"required"`
}

type CreatePurchaseReportsToBankItemRequest struct {
	IDPurchaseIgm int `json:"id_purchase_igm" validate:"required"`
}
