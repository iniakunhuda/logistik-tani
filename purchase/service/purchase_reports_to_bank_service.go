package service

import (
	purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"
	"github.com/iniakunhuda/logistik-tani/purchase/request"
	"github.com/iniakunhuda/logistik-tani/purchase/response"
)

type PurchaseReportToBankService interface {
	GenerateNoReport() (string, error)
	Create(report request.CreatePurchaseReportsToBankRequest) error
	Delete(reportId int) error
	FindById(reportId int) (response.PurchaseReportsToBankResponse, error)
	FindAll(report *purchaseigmmodel.PurchaseReportsToBank) ([]response.PurchaseReportsToBankResponse, error)
}
