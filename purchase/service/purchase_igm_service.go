package service

import (
	purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"
	"github.com/iniakunhuda/logistik-tani/purchase/request"
	"github.com/iniakunhuda/logistik-tani/purchase/response"
)

type PurchaseIgmService interface {
	Create(purchase request.CreatePurchaseIgmRequest) error
	Delete(purchaseId int) error
	FindById(purchaseId int) (response.PurchaseIgmResponse, error)
	FindAll(purchase *purchaseigmmodel.PurchaseIgm) ([]response.PurchaseIgmResponse, error)
}
