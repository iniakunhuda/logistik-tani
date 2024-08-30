package repository

import purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"

type PurchaseIgmDetailRepository interface {
	Delete(purchaseId int) error
	Update(purchase purchaseigmmodel.PurchaseIgmDetail) error
	FindById(purchaseId int) (*purchaseigmmodel.PurchaseIgmDetail, error)
	FindAll() (purchases []purchaseigmmodel.PurchaseIgmDetail, err error)
	FindLastRow() (purchase *purchaseigmmodel.PurchaseIgmDetail, err error)
	GetAllByQuery(purchase purchaseigmmodel.PurchaseIgmDetail) (purchases []purchaseigmmodel.PurchaseIgmDetail, err error)
	GetOneByQuery(purchase purchaseigmmodel.PurchaseIgmDetail) (purchaseData purchaseigmmodel.PurchaseIgmDetail, err error)
}
