package purchasereportstobank

type PurchaseReportsToBankApiListResponse struct {
	Code    int                             `json:"code"`
	Message string                          `json:"message"`
	Data    []PurchaseReportsToBankResponse `json:"data"`
}
