package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/purchase/model"
	"gorm.io/gorm"
)

type PurchaseRepositoryImpl struct {
	Db *gorm.DB
}

func NewPurchaseRepositoryImpl(Db *gorm.DB) PurchaseRepository {
	return &PurchaseRepositoryImpl{Db: Db}
}

func (t *PurchaseRepositoryImpl) FindLastRow() (purchase *model.Purchase, err error) {
	result := t.Db.Order("id desc").First(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchase, nil
}

func (t *PurchaseRepositoryImpl) Delete(purchaseId int) error {
	var purchase model.Purchase
	result := t.Db.Where("id = ?", purchaseId).Delete(&purchase)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *PurchaseRepositoryImpl) FindAll() (purchases []model.Purchase, err error) {
	var purchaseList []model.Purchase
	result := t.Db.Preload("PurchaseDetail").Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseRepositoryImpl) FindById(purchaseId int) (*model.Purchase, error) {
	var purchaseResult model.Purchase
	result := t.Db.Find(&purchaseResult, purchaseId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("purchase is not found")
	}
	return &purchaseResult, nil
}

func (t *PurchaseRepositoryImpl) Save(purchase model.Purchase, purchasesDetail []model.PurchaseDetail) error {
	// Execute the transaction
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		// Insert the Purchase record
		if err := tx.Create(&purchase).Error; err != nil {
			return err // Return error to rollback
		}

		// Insert the PurchaseDetail records
		for i := range purchasesDetail {
			purchasesDetail[i].IDPurchase = int(purchase.ID)
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

func (t *PurchaseRepositoryImpl) Update(purchase model.Purchase) error {
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

func (t *PurchaseRepositoryImpl) GetAllByQuery(purchase model.Purchase) (purchases []model.Purchase, err error) {
	var purchaseList []model.Purchase
	result := t.Db.Preload("PurchaseDetail").Where(&purchase).Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseRepositoryImpl) GetOneByQuery(purchase model.Purchase) (purchaseData model.Purchase, err error) {
	result := t.Db.Preload("PurchaseDetail").Where(&purchase).First(&purchaseData)
	if result.Error != nil {
		return model.Purchase{}, result.Error
	}
	return purchaseData, nil
}
