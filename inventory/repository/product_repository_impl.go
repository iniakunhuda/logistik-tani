package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

func (t *ProductRepositoryImpl) Save(product model.Product) error {
	result := t.Db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductRepositoryImpl) FindAll() (produks []model.Product, err error) {
	var dataList []model.Product
	result := t.Db.Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *ProductRepositoryImpl) FindByName(productName string) (*model.Product, error) {
	var dataRow model.Product
	result := t.Db.Where("name = ?", productName).First(&dataRow)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dataRow, nil
}

func (t *ProductRepositoryImpl) FindById(produkId int) (*model.Product, error) {
	var dataRow model.Product
	result := t.Db.Find(&dataRow, produkId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}
	return &dataRow, nil
}

func (t *ProductRepositoryImpl) Delete(dataId int) error {
	var data model.Product
	result := t.Db.Where("id = ?", dataId).Delete(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductRepositoryImpl) Update(produk model.Product) error {
	result := t.Db.Model(&produk).Updates(produk)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductRepositoryImpl) GetAllByQuery(produk model.Product) (produks []model.Product, err error) {
	var dataList []model.Product
	result := t.Db.Where(&produk).Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *ProductRepositoryImpl) GetOneByQuery(produk model.Product) (produkData model.Product, err error) {
	result := t.Db.Where(&produk).First(&produkData)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return produkData, nil
}
