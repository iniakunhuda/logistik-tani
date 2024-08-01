package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/sales/model"
	"gorm.io/gorm"
)

type SalesRepositoryImpl struct {
	Db *gorm.DB
}

func NewSalesRepositoryImpl(Db *gorm.DB) SalesRepository {
	return &SalesRepositoryImpl{Db: Db}
}

func (t *SalesRepositoryImpl) Delete(userId int) error {
	var user model.Sales
	result := t.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *SalesRepositoryImpl) FindAll() (produks []model.Sales, err error) {
	var produkList []model.Sales
	result := t.Db.Find(&produkList)
	if result.Error != nil {
		return nil, result.Error
	}
	return produkList, nil
}

func (t *SalesRepositoryImpl) FindById(produkId int) (*model.Sales, error) {
	var produkResult model.Sales
	result := t.Db.Find(&produkResult, produkId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("produk is not found")
	}
	return &produkResult, nil
}

func (t *SalesRepositoryImpl) Save(produk model.Sales) error {
	result := t.Db.Create(&produk)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *SalesRepositoryImpl) Update(produk model.Sales) error {
	result := t.Db.Model(&produk).Updates(produk)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *SalesRepositoryImpl) GetAllByQuery(produk model.Sales) (produks []model.Sales, err error) {
	var produkList []model.Sales
	result := t.Db.Where(&produk).Find(&produkList)
	if result.Error != nil {
		return nil, result.Error
	}
	return produkList, nil
}

func (t *SalesRepositoryImpl) GetOneByQuery(produk model.Sales) (produkData model.Sales, err error) {
	result := t.Db.Where(&produk).First(&produkData)
	if result.Error != nil {
		return model.Sales{}, result.Error
	}
	return produkData, nil
}
