package purchasereportstobank

type PurchaseReportsToBankApiDetailResponse struct {
	Code    int                           `json:"code"`
	Message string                        `json:"message"`
	Data    PurchaseReportsToBankResponse `json:"data"`
}
