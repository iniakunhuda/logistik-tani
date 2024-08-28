package service

import (
	"github.com/iniakunhuda/logistik-tani/sales/model"
	"github.com/iniakunhuda/logistik-tani/sales/request"
	"github.com/iniakunhuda/logistik-tani/sales/response"
)

type SalesIgmService interface {
	Create(sale request.CreateSalesIgmRequest) error
	Delete(saleId int) error
	FindById(saleId int, userId int) (response.SalesIgmResponse, error)
	FindAll(sale *model.SalesIgm) ([]response.SalesIgmResponse, error)
	Update(saleId int, userId int, sale request.UpdateSalesRequest) error
}
