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

func (t *SalesRepositoryImpl) FindLastRow() (sale *model.Sales, err error) {
	result := t.Db.Order("id desc").First(&sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return sale, nil
}

func (t *SalesRepositoryImpl) Delete(saleId int) error {
	var sale model.Sales
	result := t.Db.Where("id = ?", saleId).Delete(&sale)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *SalesRepositoryImpl) FindAll() (sales []model.Sales, err error) {
	var saleList []model.Sales
	result := t.Db.Preload("SalesDetail").Find(&saleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return saleList, nil
}

func (t *SalesRepositoryImpl) FindById(saleId int) (*model.Sales, error) {
	var saleResult model.Sales
	result := t.Db.Find(&saleResult, saleId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("sale is not found")
	}
	return &saleResult, nil
}

func (t *SalesRepositoryImpl) Save(sale model.Sales, salesDetail []model.SalesDetail) error {
	// Execute the transaction
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		// Insert the Sales record
		if err := tx.Create(&sale).Error; err != nil {
			return err // Return error to rollback
		}

		// Insert the SalesDetail records
		for i := range salesDetail {
			salesDetail[i].IDSales = int(sale.ID)
		}
		if err := tx.Create(&salesDetail).Error; err != nil {
			return err // Return error to rollback
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (t *SalesRepositoryImpl) Update(sale model.Sales) error {
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&sale).Updates(sale).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *SalesRepositoryImpl) GetAllByQuery(sale model.Sales) (sales []model.Sales, err error) {
	var saleList []model.Sales
	result := t.Db.Preload("SalesDetail").Where(&sale).Find(&saleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return saleList, nil
}

func (t *SalesRepositoryImpl) GetOneByQuery(sale model.Sales) (saleData model.Sales, err error) {
	result := t.Db.Preload("SalesDetail").Where(&sale).First(&saleData)
	if result.Error != nil {
		return model.Sales{}, result.Error
	}
	return saleData, nil
}
