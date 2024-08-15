package service

import (
	"github.com/iniakunhuda/logistik-tani/purchase/model"
	"github.com/iniakunhuda/logistik-tani/purchase/request"
	"github.com/iniakunhuda/logistik-tani/purchase/response"
)

type PurchaseService interface {
	Create(purchase request.CreatePurchaseRequest) error
	Update(purchaseId int, userId int, purchase request.UpdatePurchaseRequest) error
	Delete(purchaseId int) error
	FindById(purchaseId int, userId uint) (response.PurchaseResponse, error)
	FindAll(purchase *model.Purchase) ([]response.PurchaseResponse, error)
}
