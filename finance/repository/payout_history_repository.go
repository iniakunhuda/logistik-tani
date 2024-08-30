package repository

import "github.com/iniakunhuda/logistik-tani/finance/model"

type PayoutHistoryRepository interface {
	Save(purchase model.PayoutHistory) error
	Delete(purchaseId int) error
	Update(purchase model.PayoutHistory) error
	FindById(purchaseId int) (*model.PayoutHistory, error)
	FindAll() (purchases []model.PayoutHistory, err error)
	FindLastRow() (purchase *model.PayoutHistory, err error)
	GetAllByQuery(purchase model.PayoutHistory) (purchases []model.PayoutHistory, err error)
	GetOneByQuery(purchase model.PayoutHistory) (purchaseData model.PayoutHistory, err error)
}
