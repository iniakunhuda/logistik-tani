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

func (t *InventoryRepositoryImpl) Delete(userId int) error {
	var user model.Produk
	result := t.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *InventoryRepositoryImpl) FindAll() (produks []model.Produk, err error) {
	var produkList []model.Produk
	result := t.Db.Find(&produkList)
	if result.Error != nil {
		return nil, result.Error
	}
	return produkList, nil
}

func (t *InventoryRepositoryImpl) FindById(produkId int) (*model.Produk, error) {
	var produkResult model.Produk
	result := t.Db.Find(&produkResult, produkId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("Product not found")
	}
	return &produkResult, nil
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
	var produkList []model.Produk
	result := t.Db.Where(&produk).Find(&produkList)
	if result.Error != nil {
		return nil, result.Error
	}
	return produkList, nil
}

func (t *InventoryRepositoryImpl) GetOneByQuery(produk model.Produk) (produkData model.Produk, err error) {
	result := t.Db.Where(&produk).First(&produkData)
	if result.Error != nil {
		return model.Produk{}, result.Error
	}
	return produkData, nil
}
