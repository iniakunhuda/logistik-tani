package repository

import (
	"errors"

	purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"
	"gorm.io/gorm"
)

type PurchaseReportsToBankImpl struct {
	Db *gorm.DB
}

func NewPurchaseReportsToBankImpl(Db *gorm.DB) PurchaseReportsToBank {
	return &PurchaseReportsToBankImpl{Db: Db}
}

func (t *PurchaseReportsToBankImpl) FindLastRow() (purchase *purchaseigmmodel.PurchaseReportsToBank, err error) {
	result := t.Db.Preload("Details").Order("id desc").First(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchase, nil
}

func (t *PurchaseReportsToBankImpl) Delete(purchaseId int) error {
	var purchase purchaseigmmodel.PurchaseReportsToBank
	result := t.Db.Where("id = ?", purchaseId).Delete(&purchase)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *PurchaseReportsToBankImpl) FindAll() (purchases []purchaseigmmodel.PurchaseReportsToBank, err error) {
	var purchaseList []purchaseigmmodel.PurchaseReportsToBank
	result := t.Db.Preload("Details").Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseReportsToBankImpl) FindById(purchaseId int) (*purchaseigmmodel.PurchaseReportsToBank, error) {
	var purchaseResult purchaseigmmodel.PurchaseReportsToBank
	result := t.Db.Preload("Details").Find(&purchaseResult, purchaseId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("purchase is not found")
	}
	return &purchaseResult, nil
}

func (t *PurchaseReportsToBankImpl) Save(purchase purchaseigmmodel.PurchaseReportsToBank, purchasesDetail []purchaseigmmodel.PurchaseReportsToBankDetail) error {
	// Execute the transaction
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		// Insert the Purchase record
		if err := tx.Create(&purchase).Error; err != nil {
			return err // Return error to rollback
		}

		// Insert the PurchaseDetail records
		for i := range purchasesDetail {
			purchasesDetail[i].IDPurchaseReportsToBank = int(purchase.ID)
		}
		if err := tx.Create(&purchasesDetail).Error; err != nil {
			return err // Return error to rollback
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (t *PurchaseReportsToBankImpl) Update(purchase purchaseigmmodel.PurchaseReportsToBank) error {
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&purchase).Updates(purchase).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *PurchaseReportsToBankImpl) GetAllByQuery(purchase purchaseigmmodel.PurchaseReportsToBank) (purchases []purchaseigmmodel.PurchaseReportsToBank, err error) {
	var purchaseList []purchaseigmmodel.PurchaseReportsToBank
	result := t.Db.Preload("Details").Where(&purchase).Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseReportsToBankImpl) GetOneByQuery(purchase purchaseigmmodel.PurchaseReportsToBank) (purchaseData purchaseigmmodel.PurchaseReportsToBank, err error) {
	result := t.Db.Preload("Details").Where(&purchase).First(&purchaseData)
	if result.Error != nil {
		return purchaseigmmodel.PurchaseReportsToBank{}, result.Error
	}
	return purchaseData, nil
}
