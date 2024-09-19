package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	purchaseigmmodel "github.com/iniakunhuda/logistik-tani/purchase/model/purchase_igm_model"
	"github.com/iniakunhuda/logistik-tani/purchase/repository"
	"github.com/iniakunhuda/logistik-tani/purchase/repository/remote"
	"github.com/iniakunhuda/logistik-tani/purchase/request"
	"github.com/iniakunhuda/logistik-tani/purchase/response"
	userresponse "github.com/iniakunhuda/logistik-tani/purchase/response/user_response"
)

type PurchaseReportToBankServiceImpl struct {
	PurchaseReportsToBankRepository repository.PurchaseReportsToBank
	UserRemoteRepository            remote.UserRemoteRepository
	InventoryRemoteRepository       remote.InventoryRemoteRepository
	PurchaseIgmService              PurchaseIgmService
	Validate                        *validator.Validate
}

func NewPurchaseReportToBankServiceImpl(purchaseReportToBank repository.PurchaseReportsToBank, purchaseIgmService PurchaseIgmService, validate *validator.Validate) PurchaseReportToBankService {
	return &PurchaseReportToBankServiceImpl{
		PurchaseReportsToBankRepository: purchaseReportToBank,
		UserRemoteRepository:            remote.NewUserRemoteRepositoryImpl(),
		InventoryRemoteRepository:       remote.NewInventoryRemoteRepositoryImpl(""), // TODO: set bearer token
		PurchaseIgmService:              purchaseIgmService,
		Validate:                        validate,
	}
}

func (t *PurchaseReportToBankServiceImpl) GenerateNoReport() (string, error) {
	sales, err := t.PurchaseReportsToBankRepository.FindLastRow()

	lastInv := 0
	if sales != nil {
		parts := strings.Split(sales.NoReport, "-")
		if len(parts) == 2 {
			lastInv, err = strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
		}
	}

	noInv := ""
	if lastInv+1 < 10 {
		noInv = fmt.Sprintf("REPORTIGM-000%d", lastInv+1)
	} else if lastInv+1 < 100 {
		noInv = fmt.Sprintf("REPORTIGM-00%d", lastInv+1)
	} else if lastInv+1 < 1000 {
		noInv = fmt.Sprintf("REPORTIGM-0%d", lastInv+1)
	}

	return noInv, nil
}

func (t *PurchaseReportToBankServiceImpl) Create(purchase request.CreatePurchaseReportsToBankRequest) error {

	noReport, err := t.GenerateNoReport()
	if err != nil {
		return err
	}
	reportModel := purchaseigmmodel.PurchaseReportsToBank{
		NoReport:  noReport,
		DateStart: &purchase.DateStart,
		DateEnd:   &purchase.DateEnd,
		Note:      purchase.Note,
		Status:    purchase.Status,
	}

	reportDetailModel := []purchaseigmmodel.PurchaseReportsToBankDetail{}
	for _, value := range purchase.Purchases {
		reportDetailModel = append(reportDetailModel, purchaseigmmodel.PurchaseReportsToBankDetail{
			IDPurchaseIgm: value.IDPurchaseIgm,
		})
	}

	err = t.PurchaseReportsToBankRepository.Save(reportModel, reportDetailModel)
	if err != nil {
		return err
	}

	return nil
}

func (t *PurchaseReportToBankServiceImpl) Delete(purchaseId int) error {
	err := t.PurchaseReportsToBankRepository.Delete(purchaseId)
	if err != nil {
		return err
	}
	return nil
}

func (t *PurchaseReportToBankServiceImpl) FindAll(purchase *purchaseigmmodel.PurchaseReportsToBank) ([]response.PurchaseReportsToBankResponse, error) {
	result, err := t.PurchaseReportsToBankRepository.GetAllByQuery(*purchase)

	if err != nil {
		return nil, err
	}

	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return nil, err
	}
	listUser := map[uint]response.UserResponse{}
	for _, value := range users {
		listUser[value.ID] = value
	}

	// get list of products
	products, err := t.InventoryRemoteRepository.GetAll()
	if err != nil {
		return nil, err
	}
	listProducts := map[uint]response.ProductResponse{}
	for _, value := range products {
		listProducts[value.ID] = value
	}

	var purchases []response.PurchaseReportsToBankResponse
	for _, value := range result {
		purchases = append(purchases, t.formattedResponse(value))
	}

	return purchases, nil
}

func (t *PurchaseReportToBankServiceImpl) FindById(purchaseId int) (response.PurchaseReportsToBankResponse, error) {
	purchaseData, err := t.PurchaseReportsToBankRepository.GetOneByQuery(purchaseigmmodel.PurchaseReportsToBank{ID: uint(purchaseId)})
	if err != nil {
		return response.PurchaseReportsToBankResponse{}, err
	}

	// get product detail
	formatResponse := t.formattedResponse(purchaseData)
	return formatResponse, nil
}

func (t *PurchaseReportToBankServiceImpl) formattedResponse(purchaseReport purchaseigmmodel.PurchaseReportsToBank) response.PurchaseReportsToBankResponse {
	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return response.PurchaseReportsToBankResponse{}
	}
	listUser := map[int]response.UserResponse{}
	for _, value := range users {
		listUser[int(value.ID)] = value
	}

	// get user land service
	userLands, err := t.UserRemoteRepository.GetLands()
	if err != nil {
		return response.PurchaseReportsToBankResponse{}
	}
	listUserLand := map[int]userresponse.UserLandRowResponse{}
	for _, value := range userLands.Data {
		listUserLand[int(value.ID)] = value
	}

	// get list of detail
	purchasesIgm, err := t.PurchaseIgmService.FindAll(&purchaseigmmodel.PurchaseIgm{})
	if err != nil {
		return response.PurchaseReportsToBankResponse{}
	}
	listPurchasesIgm := map[int]response.PurchaseIgmResponse{}
	for _, value := range purchasesIgm {
		listPurchasesIgm[int(value.ID)] = value
	}

	// detail report to bank
	detail := []response.PurchaseReportsToBankDetailResponse{}
	for _, value := range purchaseReport.Details {
		detail = append(detail, response.PurchaseReportsToBankDetailResponse{
			PurchaseIgmResponse: listPurchasesIgm[int(value.IDPurchaseIgm)],
		})
	}

	jcart, _ := json.Marshal(purchaseReport)
	fmt.Println(string(jcart))

	return response.PurchaseReportsToBankResponse{
		ID:        purchaseReport.ID,
		Note:      purchaseReport.Note,
		Status:    purchaseReport.Status,
		NoReport:  purchaseReport.NoReport,
		DateStart: purchaseReport.DateStart,
		DateEnd:   purchaseReport.DateEnd,
		Detail:    detail,
	}
}
