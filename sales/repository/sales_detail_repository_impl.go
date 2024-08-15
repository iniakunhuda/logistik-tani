package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/sales/model"
	"gorm.io/gorm"
)

type SalesDetailRepositoryImpl struct {
	Db *gorm.DB
}

func NewSalesDetailRepositoryImpl(Db *gorm.DB) SalesDetailRepository {
	return &SalesDetailRepositoryImpl{Db: Db}
}

func (t *SalesDetailRepositoryImpl) FindLastRow() (sale *model.SalesDetail, err error) {
	result := t.Db.Order("id desc").First(&sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return sale, nil
}

func (t *SalesDetailRepositoryImpl) Delete(saleId int) error {
	var sale model.SalesDetail
	result := t.Db.Where("id = ?", saleId).Delete(&sale)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *SalesDetailRepositoryImpl) FindAll() (sales []model.SalesDetail, err error) {
	var saleList []model.SalesDetail
	result := t.Db.Find(&saleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return saleList, nil
}

func (t *SalesDetailRepositoryImpl) FindById(saleId int) (*model.SalesDetail, error) {
	var saleResult model.SalesDetail
	result := t.Db.Find(&saleResult, saleId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("sale is not found")
	}
	return &saleResult, nil
}

func (t *SalesDetailRepositoryImpl) Save(sale model.SalesDetail) error {
	result := t.Db.Create(&sale)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *SalesDetailRepositoryImpl) Update(sale model.SalesDetail) error {
	result := t.Db.Model(&sale).Updates(sale)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *SalesDetailRepositoryImpl) GetAllByQuery(sale model.SalesDetail) (sales []model.SalesDetail, err error) {
	var saleList []model.SalesDetail
	result := t.Db.Where(&sale).Find(&saleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return saleList, nil
}

func (t *SalesDetailRepositoryImpl) GetOneByQuery(sale model.SalesDetail) (saleData model.SalesDetail, err error) {
	result := t.Db.Where(&sale).First(&saleData)
	if result.Error != nil {
		return model.SalesDetail{}, result.Error
	}
	return saleData, nil
}
