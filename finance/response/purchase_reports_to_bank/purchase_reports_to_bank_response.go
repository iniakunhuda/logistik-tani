package purchasereportstobank

import "time"

type PurchaseReportsToBankResponse struct {
	ID        uint                                  `json:"id"`
	NoReport  string                                `json:"no_report"`
	DateStart *time.Time                            `json:"date_start"`
	DateEnd   *time.Time                            `json:"date_end"`
	Note      string                                `json:"note"`
	Status    string                                `json:"status"`
	Detail    []PurchaseReportsToBankDetailResponse `json:"detail"`
}

type PurchaseReportsToBankDetailResponse struct {
	PurchaseIgmResponse
}
