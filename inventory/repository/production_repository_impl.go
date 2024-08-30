package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"gorm.io/gorm"
)

type ProductionRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductionRepositoryImpl(Db *gorm.DB) ProductionRepository {
	return &ProductionRepositoryImpl{Db: Db}
}

func (t *ProductionRepositoryImpl) Save(panen model.Production) error {
	result := t.Db.Create(&panen)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductionRepositoryImpl) FindAll() (panens []model.Production, err error) {
	var dataList []model.Production
	result := t.Db.Preload("Histories").Preload("Histories.ProductOwner").Preload("Histories.ProductOwner.Product").Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *ProductionRepositoryImpl) FindById(panenId int) (*model.Production, error) {
	var dataRow model.Production
	result := t.Db.Preload("Histories").Preload("Histories.ProductOwner").Preload("Histories.ProductOwner.Product").Find(&dataRow, panenId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("panen not found")
	}
	return &dataRow, nil
}

func (t *ProductionRepositoryImpl) Delete(dataId int) error {
	var data model.Production
	result := t.Db.Where("id = ?", dataId).Delete(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductionRepositoryImpl) Update(panen model.Production) error {
	result := t.Db.Model(&panen).Updates(panen)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductionRepositoryImpl) GetAllByQuery(panen model.Production) (panens []model.Production, err error) {
	var dataList []model.Production
	result := t.Db.Where(&panen).Preload("Histories").Preload("Histories.ProductOwner").Preload("Histories.ProductOwner.Product").Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *ProductionRepositoryImpl) GetOneByQuery(panen model.Production) (panenData model.Production, err error) {
	result := t.Db.Where(&panen).Preload("Histories").Preload("Histories.ProductOwner").Preload("Histories.ProductOwner.Product").First(&panenData)
	if result.Error != nil {
		return model.Production{}, result.Error
	}
	return panenData, nil
}
