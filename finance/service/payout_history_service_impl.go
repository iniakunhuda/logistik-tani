package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/finance/model"
	"github.com/iniakunhuda/logistik-tani/finance/repository"
	"github.com/iniakunhuda/logistik-tani/finance/repository/remote"
	"github.com/iniakunhuda/logistik-tani/finance/request"
	"github.com/iniakunhuda/logistik-tani/finance/response"
)

type PayoutHistoryServiceImpl struct {
	TokenAuth            string
	PayoutRepository     repository.PayoutHistoryRepository
	UserRemoteRepository remote.UserRemoteRepository
	Validate             *validator.Validate
}

func NewPayoutHistoryServiceImpl(purchasesRepository repository.PayoutHistoryRepository, validate *validator.Validate) PayoutHistoryService {
	return &PayoutHistoryServiceImpl{
		PayoutRepository:     purchasesRepository,
		UserRemoteRepository: remote.NewUserRemoteRepositoryImpl(),
		Validate:             validate,
	}
}

func (t *PayoutHistoryServiceImpl) GenerateNoInvoice() (string, error) {
	sales, err := t.PayoutRepository.FindLastRow()

	lastInv := 0
	if sales != nil {
		parts := strings.Split(sales.NoInvoice, "-")
		if len(parts) == 2 {
			lastInv, err = strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
		}
	}

	noInv := ""
	if lastInv+1 < 10 {
		noInv = fmt.Sprintf("PAYOUT-000%d", lastInv+1)
	} else if lastInv+1 < 100 {
		noInv = fmt.Sprintf("PAYOUT-00%d", lastInv+1)
	} else if lastInv+1 < 1000 {
		noInv = fmt.Sprintf("PAYOUT-0%d", lastInv+1)
	}

	return noInv, nil
}

func (t *PayoutHistoryServiceImpl) Create(purchase request.CreatePayoutRequest) error {

	noInv, err := t.GenerateNoInvoice()
	if err != nil {
		return err
	}
	purchaseModel := model.PayoutHistory{
		NoInvoice:               noInv,
		IDSender:                purchase.IDSender,
		IDReceiver:              purchase.IDReceiver,
		IDPurchaseReportsToBank: purchase.IDPurchaseReportsToBank,
		TotalAmount:             purchase.TotalAmount,
		BankNote:                purchase.BankNote,
		DatePayout:              purchase.DatePayout,
		Status:                  purchase.Status,
		CreatedDate:             purchase.CreatedDate,
	}

	err = t.PayoutRepository.Save(purchaseModel)
	if err != nil {
		return err
	}

	return nil
}

func (t *PayoutHistoryServiceImpl) Update(purchaseId int, purchases request.UpdatePayoutRequest) error {
	// Validate the request
	payoutDb, err := t.PayoutRepository.GetOneByQuery(model.PayoutHistory{ID: uint(purchaseId)})
	if err != nil {
		return err
	}

	// Check if status already approved or rejected
	if payoutDb.Status == "approved" {
		return errors.New("error! transaksi sudah diapprove")
	}

	statusArr := []string{"pending", "approved", "rejected"}
	isValid := false
	for _, value := range statusArr {
		if value == purchases.Status {
			isValid = true
			break
		}
	}

	if !isValid {
		return errors.New("error! status tidak valid (pending, approved, rejected)")
	}

	// time today
	today := time.Now()

	filter := model.PayoutHistory{}
	if purchases.Status == "approved" {
		filter = model.PayoutHistory{
			ID:              uint(purchaseId),
			ApprovedDate:    &today,
			ApprovedMessage: &purchases.Message,
			Status:          purchases.Status,
		}

		// trigger update saldo user
		err = t.UserRemoteRepository.AddSaldo(strconv.Itoa(payoutDb.IDReceiver), int(payoutDb.TotalAmount))
		if err != nil {
			return err
		}
	}

	if purchases.Status == "rejected" {
		filter = model.PayoutHistory{
			ID:              uint(purchaseId),
			RejectedDate:    &today,
			RejectedMessage: &purchases.Message,
			Status:          purchases.Status,
		}
	}

	// Update the purchase status
	err = t.PayoutRepository.Update(filter)
	if err != nil {
		return err
	}

	return nil
}

func (t *PayoutHistoryServiceImpl) Delete(purchaseId int) error {
	err := t.PayoutRepository.Delete(purchaseId)
	if err != nil {
		return err
	}
	return nil
}

func (t *PayoutHistoryServiceImpl) FindAll(purchase *model.PayoutHistory) ([]response.PayoutHistoryResponse, error) {
	result, err := t.PayoutRepository.GetAllByQuery(*purchase)

	if err != nil {
		return nil, err
	}
	var purchases []response.PayoutHistoryResponse
	for _, value := range result {
		purchases = append(purchases, t.formattedResponse(value))
	}

	return purchases, nil
}

func (t *PayoutHistoryServiceImpl) FindById(purchaseId int) (response.PayoutHistoryResponse, error) {
	purchaseData, err := t.PayoutRepository.GetOneByQuery(model.PayoutHistory{ID: uint(purchaseId)})
	if err != nil {
		return response.PayoutHistoryResponse{}, err
	}

	formatResponse := t.formattedResponse(purchaseData)
	return formatResponse, nil
}

func (t *PayoutHistoryServiceImpl) formattedResponse(value model.PayoutHistory) response.PayoutHistoryResponse {
	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return response.PayoutHistoryResponse{}
	}
	listUser := map[int]response.UserResponse{}
	for _, value := range users {
		listUser[int(value.ID)] = value
	}

	return response.PayoutHistoryResponse{
		IDSender:                value.IDSender,
		IDReceiver:              value.IDReceiver,
		NoInvoice:               value.NoInvoice,
		TotalAmount:             value.TotalAmount,
		BankNote:                value.BankNote,
		DatePayout:              value.DatePayout,
		IDPurchaseReportsToBank: value.IDPurchaseReportsToBank,
		Status:                  value.Status,
		CreatedDate:             value.CreatedDate,
		ApprovedDate:            value.ApprovedDate,
		ApprovedMessage:         value.ApprovedMessage,
		RejectedDate:            value.RejectedDate,
		RejectedMessage:         value.RejectedMessage,

		SenderDetail:   listUser[value.IDSender],
		ReceiverDetail: listUser[value.IDReceiver],
	}
}
