package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"gorm.io/gorm"
)

type ProductOwnerRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductOwnerRepositoryImpl(Db *gorm.DB) ProductOwnerRepository {
	return &ProductOwnerRepositoryImpl{Db: Db}
}

func (t *ProductOwnerRepositoryImpl) Save(product model.ProductOwner) error {
	result := t.Db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductOwnerRepositoryImpl) FindAll() (produks []model.ProductOwner, err error) {
	var dataList []model.ProductOwner
	result := t.Db.Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *ProductOwnerRepositoryImpl) FindById(produkId int) (*model.ProductOwner, error) {
	var dataRow model.ProductOwner
	result := t.Db.Find(&dataRow, produkId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("ProductOwner not found")
	}
	return &dataRow, nil
}

func (t *ProductOwnerRepositoryImpl) Delete(dataId int) error {
	var data model.ProductOwner
	result := t.Db.Where("id = ?", dataId).Delete(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductOwnerRepositoryImpl) Update(produk model.ProductOwner) error {
	result := t.Db.Model(&produk).Updates(produk)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *ProductOwnerRepositoryImpl) GetAllByQuery(produk model.ProductOwner) (produks []model.ProductOwner, err error) {
	var dataList []model.ProductOwner
	result := t.Db.Where(&produk).Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *ProductOwnerRepositoryImpl) GetOneByQuery(produk model.ProductOwner) (produkData model.ProductOwner, err error) {
	result := t.Db.Where(&produk).First(&produkData)
	if result.Error != nil {
		return model.ProductOwner{}, result.Error
	}
	return produkData, nil
}

// move product to filter class,
func (t *ProductOwnerRepositoryImpl) GetAllByProduk(product model.Product, userId ...string) (products []model.ProductOwner, err error) {
	var dataList []model.ProductOwner

	result := t.Db.Preload("Product").Joins("JOIN product ON product.id = product_owner.id_product")
	if len(userId) > 0 {
		result = result.Where("product_owner.id_user = ?", userId[0])
	}

	if product.Name != "" {
		result = result.Where("product.name = ?", product.Name)
	}

	if product.Category != "" {
		result = result.Where("product.category = ?", product.Category)
	}

	if product.ID != 0 {
		result = result.Where("product_owner.id = ?", product.ID)
	}

	result = result.Find(&dataList)

	if result.Error != nil {
		return nil, result.Error
	}

	return dataList, nil
}
