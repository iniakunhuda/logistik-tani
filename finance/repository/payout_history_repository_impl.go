package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/finance/model"
	"gorm.io/gorm"
)

type PayoutHistoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewPayoutHistoryRepositoryImpl(Db *gorm.DB) PayoutHistoryRepository {
	return &PayoutHistoryRepositoryImpl{Db: Db}
}

func (t *PayoutHistoryRepositoryImpl) FindLastRow() (payout *model.PayoutHistory, err error) {
	result := t.Db.Order("id desc").First(&payout)
	if result.Error != nil {
		return nil, result.Error
	}
	return payout, nil
}

func (t *PayoutHistoryRepositoryImpl) Delete(payoutId int) error {
	var payout model.PayoutHistory
	result := t.Db.Where("id = ?", payoutId).Delete(&payout)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *PayoutHistoryRepositoryImpl) FindAll() (payouts []model.PayoutHistory, err error) {
	var payoutList []model.PayoutHistory
	result := t.Db.Find(&payoutList)
	if result.Error != nil {
		return nil, result.Error
	}
	return payoutList, nil
}

func (t *PayoutHistoryRepositoryImpl) FindById(payoutId int) (*model.PayoutHistory, error) {
	var payoutResult model.PayoutHistory
	result := t.Db.Find(&payoutResult, payoutId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("payout is not found")
	}
	return &payoutResult, nil
}

func (t *PayoutHistoryRepositoryImpl) Save(payout model.PayoutHistory) error {
	// Execute the transaction
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		// Insert the Purchase record
		if err := tx.Create(&payout).Error; err != nil {
			return err // Return error to rollback
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (t *PayoutHistoryRepositoryImpl) Update(payout model.PayoutHistory) error {
	err := t.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&payout).Updates(payout).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *PayoutHistoryRepositoryImpl) GetAllByQuery(payout model.PayoutHistory) (payouts []model.PayoutHistory, err error) {
	var payoutList []model.PayoutHistory
	result := t.Db.Where(&payout).Find(&payoutList)
	if result.Error != nil {
		return nil, result.Error
	}
	return payoutList, nil
}

func (t *PayoutHistoryRepositoryImpl) GetOneByQuery(payout model.PayoutHistory) (payoutData model.PayoutHistory, err error) {
	result := t.Db.Where(&payout).First(&payoutData)
	if result.Error != nil {
		return model.PayoutHistory{}, result.Error
	}
	return payoutData, nil
}
