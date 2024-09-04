package repository

import purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"

type PurchaseReportsToBank interface {
	Save(purchase purchaseigmmodel.PurchaseReportsToBank, purchasesDetail []purchaseigmmodel.PurchaseReportsToBankDetail) error
	Delete(purchaseId int) error
	Update(purchase purchaseigmmodel.PurchaseReportsToBank) error
	FindById(purchaseId int) (*purchaseigmmodel.PurchaseReportsToBank, error)
	FindAll() (purchases []purchaseigmmodel.PurchaseReportsToBank, err error)
	FindLastRow() (purchase *purchaseigmmodel.PurchaseReportsToBank, err error)
	GetAllByQuery(purchase purchaseigmmodel.PurchaseReportsToBank) (purchases []purchaseigmmodel.PurchaseReportsToBank, err error)
	GetOneByQuery(purchase purchaseigmmodel.PurchaseReportsToBank) (purchaseData purchaseigmmodel.PurchaseReportsToBank, err error)
}
