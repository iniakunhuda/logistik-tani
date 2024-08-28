package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/sales/model"
	"gorm.io/gorm"
)

type SalesIgmRepositoryImpl struct {
	Db *gorm.DB
}

func NewSalesIgmRepositoryImpl(Db *gorm.DB) SalesIgmRepository {
	return &SalesIgmRepositoryImpl{Db: Db}
}

func (t *SalesIgmRepositoryImpl) FindLastRow() (sale *model.SalesIgm, err error) {
	result := t.Db.Order("id desc").First(&sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return sale, nil
}

func (t *SalesIgmRepositoryImpl) Delete(saleId int) error {
	var sale model.SalesIgm
	result := t.Db.Where("id = ?", saleId).Delete(&sale)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *SalesIgmRepositoryImpl) FindAll() (sales []model.SalesIgm, err error) {
	var saleList []model.SalesIgm
	result := t.Db.Find(&saleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return saleList, nil
}

func (t *SalesIgmRepositoryImpl) FindById(saleId int) (*model.SalesIgm, error) {
	var saleResult model.SalesIgm
	result := t.Db.Find(&saleResult, saleId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("sale is not found")
	}
	return &saleResult, nil
}

func (t *SalesIgmRepositoryImpl) Save(sale model.SalesIgm) error {
	// Execute the transaction
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		// Insert the Sales record
		if err := tx.Create(&sale).Error; err != nil {
			return err // Return error to rollback
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (t *SalesIgmRepositoryImpl) Update(sale model.SalesIgm) error {
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

func (t *SalesIgmRepositoryImpl) GetAllByQuery(sale model.SalesIgm) (sales []model.SalesIgm, err error) {
	var saleList []model.SalesIgm
	result := t.Db.Where(&sale).Find(&saleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return saleList, nil
}

func (t *SalesIgmRepositoryImpl) GetOneByQuery(sale model.SalesIgm) (saleData model.SalesIgm, err error) {
	result := t.Db.Where(&sale).First(&saleData)
	if result.Error != nil {
		return model.SalesIgm{}, result.Error
	}
	return saleData, nil
}
