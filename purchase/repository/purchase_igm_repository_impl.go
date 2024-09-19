package repository

import (
	"errors"

	purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"
	"gorm.io/gorm"
)

type PurchaseIgmRepositoryImpl struct {
	Db *gorm.DB
}

func NewPurchaseIgmRepositoryImpl(Db *gorm.DB) PurchaseIgmRepository {
	return &PurchaseIgmRepositoryImpl{Db: Db}
}

func (t *PurchaseIgmRepositoryImpl) FindLastRow() (purchase *purchaseigmmodel.PurchaseIgm, err error) {
	result := t.Db.Order("id desc").First(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchase, nil
}

func (t *PurchaseIgmRepositoryImpl) Delete(purchaseId int) error {
	var purchase purchaseigmmodel.PurchaseIgm
	result := t.Db.Where("id = ?", purchaseId).Delete(&purchase)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *PurchaseIgmRepositoryImpl) FindAll() (purchases []purchaseigmmodel.PurchaseIgm, err error) {
	var purchaseList []purchaseigmmodel.PurchaseIgm
	result := t.Db.Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseIgmRepositoryImpl) FindById(purchaseId int) (*purchaseigmmodel.PurchaseIgm, error) {
	var purchaseResult purchaseigmmodel.PurchaseIgm
	result := t.Db.Find(&purchaseResult, purchaseId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("purchase is not found")
	}
	return &purchaseResult, nil
}

func (t *PurchaseIgmRepositoryImpl) Save(purchase purchaseigmmodel.PurchaseIgm, purchasesDetail []purchaseigmmodel.PurchaseIgmDetail) error {
	// Execute the transaction
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		// Insert the Purchase record
		if err := tx.Create(&purchase).Error; err != nil {
			return err // Return error to rollback
		}

		// Insert the PurchaseDetail records
		for i := range purchasesDetail {
			purchasesDetail[i].IDPurchaseIgm = int(purchase.ID)
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

func (t *PurchaseIgmRepositoryImpl) Update(purchase purchaseigmmodel.PurchaseIgm) error {
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

func (t *PurchaseIgmRepositoryImpl) GetAllByQuery(purchase purchaseigmmodel.PurchaseIgm) (purchases []purchaseigmmodel.PurchaseIgm, err error) {
	var purchaseList []purchaseigmmodel.PurchaseIgm
	result := t.Db.Where(&purchase).Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseIgmRepositoryImpl) GetOneByQuery(purchase purchaseigmmodel.PurchaseIgm) (purchaseData purchaseigmmodel.PurchaseIgm, err error) {
	result := t.Db.Where(&purchase).First(&purchaseData)
	if result.Error != nil {
		return purchaseigmmodel.PurchaseIgm{}, result.Error
	}
	return purchaseData, nil
}
