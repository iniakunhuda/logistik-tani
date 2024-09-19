package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"gorm.io/gorm"
)

type ProductionDetailRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductionDetailRepositoryImpl(Db *gorm.DB) ProductionDetailRepository {
	return &ProductionDetailRepositoryImpl{Db: Db}
}

func (t *ProductionDetailRepositoryImpl) Save(panen model.ProductionDetail) error {
	result := t.Db.Create(&panen)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductionDetailRepositoryImpl) FindAll() (panens []model.ProductionDetail, err error) {
	var dataList []model.ProductionDetail
	result := t.Db.Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *ProductionDetailRepositoryImpl) FindById(panenId int) (*model.ProductionDetail, error) {
	var dataRow model.ProductionDetail
	result := t.Db.Find(&dataRow, panenId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("panen not found")
	}
	return &dataRow, nil
}

func (t *ProductionDetailRepositoryImpl) Delete(dataId int) error {
	var data model.ProductionDetail
	result := t.Db.Where("id = ?", dataId).Delete(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductionDetailRepositoryImpl) Update(panen model.ProductionDetail) error {
	result := t.Db.Model(&panen).Updates(panen)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductionDetailRepositoryImpl) GetAllByQuery(panen model.ProductionDetail) (panens []model.ProductionDetail, err error) {
	var dataList []model.ProductionDetail
	result := t.Db.Where(&panen).Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *ProductionDetailRepositoryImpl) GetOneByQuery(panen model.ProductionDetail) (panenData model.ProductionDetail, err error) {
	result := t.Db.Where(&panen).First(&panenData)
	if result.Error != nil {
		return model.ProductionDetail{}, result.Error
	}
	return panenData, nil
}
