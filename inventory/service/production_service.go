package service

import (
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type ProductionService interface {
	Create(production request.CreateProductionRequest) error
	Delete(productionId int) error
	FindById(productionId int) (response.ProductionResponse, error)
	FindAll(production *model.Production) ([]response.ProductionResponse, error)
	Update(productionId int, production request.UpdateProductionRequest) error

	CreateRiwayat(production request.CreateProductionDetailRequest) error
}
