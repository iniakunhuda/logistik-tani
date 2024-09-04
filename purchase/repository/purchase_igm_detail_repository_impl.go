package repository

import (
	"errors"

	purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"
	"gorm.io/gorm"
)

type PurchaseIgmDetailRepositoryImpl struct {
	Db *gorm.DB
}

func NewPurchaseIgmDetailRepositoryImpl(Db *gorm.DB) PurchaseIgmDetailRepository {
	return &PurchaseIgmDetailRepositoryImpl{Db: Db}
}

func (t *PurchaseIgmDetailRepositoryImpl) FindLastRow() (purchase *purchaseigmmodel.PurchaseIgmDetail, err error) {
	result := t.Db.Order("id desc").First(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchase, nil
}

func (t *PurchaseIgmDetailRepositoryImpl) Delete(purchaseId int) error {
	var purchase purchaseigmmodel.PurchaseIgmDetail
	result := t.Db.Where("id = ?", purchaseId).Delete(&purchase)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *PurchaseIgmDetailRepositoryImpl) FindAll() (purchases []purchaseigmmodel.PurchaseIgmDetail, err error) {
	var purchaseList []purchaseigmmodel.PurchaseIgmDetail
	result := t.Db.Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseIgmDetailRepositoryImpl) FindById(purchaseId int) (*purchaseigmmodel.PurchaseIgmDetail, error) {
	var purchaseResult purchaseigmmodel.PurchaseIgmDetail
	result := t.Db.Find(&purchaseResult, purchaseId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("purchase is not found")
	}
	return &purchaseResult, nil
}

func (t *PurchaseIgmDetailRepositoryImpl) Update(purchase purchaseigmmodel.PurchaseIgmDetail) error {
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

func (t *PurchaseIgmDetailRepositoryImpl) GetAllByQuery(purchase purchaseigmmodel.PurchaseIgmDetail) (purchases []purchaseigmmodel.PurchaseIgmDetail, err error) {
	var purchaseList []purchaseigmmodel.PurchaseIgmDetail
	result := t.Db.Where(&purchase).Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseIgmDetailRepositoryImpl) GetOneByQuery(purchase purchaseigmmodel.PurchaseIgmDetail) (purchaseData purchaseigmmodel.PurchaseIgmDetail, err error) {
	result := t.Db.Where(&purchase).First(&purchaseData)
	if result.Error != nil {
		return purchaseigmmodel.PurchaseIgmDetail{}, result.Error
	}
	return purchaseData, nil
}
