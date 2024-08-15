package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"gorm.io/gorm"
)

type InventoryPetaniRepositoryImpl struct {
	Db *gorm.DB
}

func NewInventoryPetaniRepositoryImpl(Db *gorm.DB) InventoryPetaniRepository {
	return &InventoryPetaniRepositoryImpl{Db: Db}
}

func (t *InventoryPetaniRepositoryImpl) Delete(dataId int) error {
	var data model.ProdukPetani
	result := t.Db.Where("id = ?", dataId).Delete(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *InventoryPetaniRepositoryImpl) FindAll() (produks []model.ProdukPetani, err error) {
	var dataList []model.ProdukPetani
	result := t.Db.Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *InventoryPetaniRepositoryImpl) FindById(produkId int) (*model.ProdukPetani, error) {
	var dataRow model.ProdukPetani
	result := t.Db.Find(&dataRow, produkId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("Product not found")
	}
	return &dataRow, nil
}

func (t *InventoryPetaniRepositoryImpl) Save(produk model.ProdukPetani) error {
	result := t.Db.Create(&produk)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *InventoryPetaniRepositoryImpl) Update(produk model.ProdukPetani) error {
	result := t.Db.Model(&produk).Updates(produk)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *InventoryPetaniRepositoryImpl) GetAllByQuery(produk model.ProdukPetani) (produks []model.ProdukPetani, err error) {
	var dataList []model.ProdukPetani
	result := t.Db.Where(&produk).Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *InventoryPetaniRepositoryImpl) GetOneByQuery(produk model.ProdukPetani) (produkData model.ProdukPetani, err error) {
	result := t.Db.Where(&produk).First(&produkData)
	if result.Error != nil {
		return model.ProdukPetani{}, result.Error
	}
	return produkData, nil
}
