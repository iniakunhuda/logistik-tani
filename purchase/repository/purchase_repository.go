package repository

import "github.com/iniakunhuda/logistik-tani/purchase/model"

type PurchaseRepository interface {
	Save(purchase model.Purchase, purchasesDetail []model.PurchaseDetail) error
	Delete(purchaseId int) error
	Update(purchase model.Purchase) error
	FindById(purchaseId int) (*model.Purchase, error)
	FindAll() (purchases []model.Purchase, err error)
	FindLastRow() (purchase *model.Purchase, err error)
	GetAllByQuery(purchase model.Purchase) (purchases []model.Purchase, err error)
	GetOneByQuery(purchase model.Purchase) (purchaseData model.Purchase, err error)
}
