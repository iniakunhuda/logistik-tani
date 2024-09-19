package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/purchase/model"
	"gorm.io/gorm"
)

type PurchaseDetailRepositoryImpl struct {
	Db *gorm.DB
}

func NewPurchaseDetailRepositoryImpl(Db *gorm.DB) PurchaseDetailRepository {
	return &PurchaseDetailRepositoryImpl{Db: Db}
}

func (t *PurchaseDetailRepositoryImpl) FindLastRow() (purchase *model.PurchaseDetail, err error) {
	result := t.Db.Order("id desc").First(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchase, nil
}

func (t *PurchaseDetailRepositoryImpl) Delete(purchaseId int) error {
	var purchase model.PurchaseDetail
	result := t.Db.Where("id = ?", purchaseId).Delete(&purchase)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *PurchaseDetailRepositoryImpl) FindAll() (purchases []model.PurchaseDetail, err error) {
	var purchaseList []model.PurchaseDetail
	result := t.Db.Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseDetailRepositoryImpl) FindById(purchaseId int) (*model.PurchaseDetail, error) {
	var purchaseResult model.PurchaseDetail
	result := t.Db.Find(&purchaseResult, purchaseId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("purchase is not found")
	}
	return &purchaseResult, nil
}

func (t *PurchaseDetailRepositoryImpl) Save(purchase model.PurchaseDetail) error {
	result := t.Db.Create(&purchase)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *PurchaseDetailRepositoryImpl) Update(purchase model.PurchaseDetail) error {
	result := t.Db.Model(&purchase).Updates(purchase)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *PurchaseDetailRepositoryImpl) GetAllByQuery(purchase model.PurchaseDetail) (purchases []model.PurchaseDetail, err error) {
	var purchaseList []model.PurchaseDetail
	result := t.Db.Where(&purchase).Find(&purchaseList)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchaseList, nil
}

func (t *PurchaseDetailRepositoryImpl) GetOneByQuery(purchase model.PurchaseDetail) (purchaseData model.PurchaseDetail, err error) {
	result := t.Db.Where(&purchase).First(&purchaseData)
	if result.Error != nil {
		return model.PurchaseDetail{}, result.Error
	}
	return purchaseData, nil
}
