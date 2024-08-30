package service

import (
	"github.com/iniakunhuda/logistik-tani/finance/model"
	"github.com/iniakunhuda/logistik-tani/finance/request"
	"github.com/iniakunhuda/logistik-tani/finance/response"
)

type PayoutHistoryService interface {
	Create(purchase request.CreatePayoutRequest) error
	Update(purchaseId int, purchase request.UpdatePayoutRequest) error
	Delete(purchaseId int) error
	FindById(purchaseId int) (response.PayoutHistoryResponse, error)
	FindAll(purchase *model.PayoutHistory) ([]response.PayoutHistoryResponse, error)
}
