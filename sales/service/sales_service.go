package service

import (
	"github.com/iniakunhuda/logistik-tani/sales/request"
	"github.com/iniakunhuda/logistik-tani/sales/response"
)

type SalesService interface {
	Create(sale request.CreateSalesRequest) error
	Delete(saleId int) error
	FindById(saleId int) (response.SalesResponse, error)
	FindAll() ([]response.SalesResponse, error)
}
