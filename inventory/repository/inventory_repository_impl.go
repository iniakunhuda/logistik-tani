package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"gorm.io/gorm"
)

type InventoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewInventoryRepositoryImpl(Db *gorm.DB) InventoryRepository {
	return &InventoryRepositoryImpl{Db: Db}
}

func (t *InventoryRepositoryImpl) Delete(dataId int) error {
	var data model.Produk
	result := t.Db.Where("id = ?", dataId).Delete(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *InventoryRepositoryImpl) FindAll() (produks []model.Produk, err error) {
	var dataList []model.Produk
	result := t.Db.Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *InventoryRepositoryImpl) FindById(produkId int) (*model.Produk, error) {
	var dataRow model.Produk
	result := t.Db.Find(&dataRow, produkId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("Product not found")
	}
	return &dataRow, nil
}

func (t *InventoryRepositoryImpl) Save(produk model.Produk) error {
	result := t.Db.Create(&produk)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *InventoryRepositoryImpl) Update(produk model.Produk) error {
	result := t.Db.Model(&produk).Updates(produk)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *InventoryRepositoryImpl) GetAllByQuery(produk model.Produk) (produks []model.Produk, err error) {
	var dataList []model.Produk
	result := t.Db.Where(&produk).Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *InventoryRepositoryImpl) GetOneByQuery(produk model.Produk) (produkData model.Produk, err error) {
	result := t.Db.Where(&produk).First(&produkData)
	if result.Error != nil {
		return model.Produk{}, result.Error
	}
	return produkData, nil
}
