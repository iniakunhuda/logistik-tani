package repository

import purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"

type PurchaseIgmRepository interface {
	Save(purchase purchaseigmmodel.PurchaseIgm, purchasesDetail []purchaseigmmodel.PurchaseIgmDetail) error
	Delete(purchaseId int) error
	Update(purchase purchaseigmmodel.PurchaseIgm) error
	FindById(purchaseId int) (*purchaseigmmodel.PurchaseIgm, error)
	FindAll() (purchases []purchaseigmmodel.PurchaseIgm, err error)
	FindLastRow() (purchase *purchaseigmmodel.PurchaseIgm, err error)
	GetAllByQuery(purchase purchaseigmmodel.PurchaseIgm) (purchases []purchaseigmmodel.PurchaseIgm, err error)
	GetOneByQuery(purchase purchaseigmmodel.PurchaseIgm) (purchaseData purchaseigmmodel.PurchaseIgm, err error)
}
