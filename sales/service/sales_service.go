package service

import (
	"github.com/iniakunhuda/logistik-tani/sales/model"
	"github.com/iniakunhuda/logistik-tani/sales/request"
	"github.com/iniakunhuda/logistik-tani/sales/response"
)

type SalesService interface {
	Create(sale request.CreateSalesRequest) error
	Delete(saleId int) error
	FindById(saleId int, userId uint) (response.SalesResponse, error)
	FindAll(sale *model.Sales) ([]response.SalesResponse, error)
}
