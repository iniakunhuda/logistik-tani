package remote

import (
	purchasereportstobank "github.com/iniakunhuda/logistik-tani/finance/response/purchase_reports_to_bank"
)

type PurchaseReportsToBankRepository interface {
	GetAll() ([]purchasereportstobank.PurchaseReportsToBankResponse, error)
	Find(id string) (purchasereportstobank.PurchaseReportsToBankResponse, error)
}
