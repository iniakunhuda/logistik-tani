package service

import (
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

type PurchaseIgmServiceImpl struct {
	PurchaseIgmRepository       repository.PurchaseIgmRepository
	PurchaseIgmDetailRepository repository.PurchaseIgmDetailRepository
	UserRemoteRepository        remote.UserRemoteRepository
	InventoryRemoteRepository   remote.InventoryRemoteRepository
	Validate                    *validator.Validate
}

func NewPurchaseIgmServiceImpl(purchaseIgmRepository repository.PurchaseIgmRepository, purchaseIgmDetailRepository repository.PurchaseIgmDetailRepository, validate *validator.Validate) PurchaseIgmService {
	return &PurchaseIgmServiceImpl{
		PurchaseIgmRepository:       purchaseIgmRepository,
		PurchaseIgmDetailRepository: purchaseIgmDetailRepository,
		UserRemoteRepository:        remote.NewUserRemoteRepositoryImpl(),
		InventoryRemoteRepository:   remote.NewInventoryRemoteRepositoryImpl(""), // TODO: set bearer token
		Validate:                    validate,
	}
}

func (t *PurchaseIgmServiceImpl) GenerateNoInvoice() (string, error) {
	sales, err := t.PurchaseIgmRepository.FindLastRow()

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
		noInv = fmt.Sprintf("PURCHASEIGM-000%d", lastInv+1)
	} else if lastInv+1 < 100 {
		noInv = fmt.Sprintf("PURCHASEIGM-00%d", lastInv+1)
	} else if lastInv+1 < 1000 {
		noInv = fmt.Sprintf("PURCHASEIGM-0%d", lastInv+1)
	}

	return noInv, nil
}

func (t *PurchaseIgmServiceImpl) Create(purchase request.CreatePurchaseIgmRequest) error {

	noInv, err := t.GenerateNoInvoice()
	if err != nil {
		return err
	}
	totalHargaPurchase := 0.0
	purchaseModel := purchaseigmmodel.PurchaseIgm{
		NoInvoice:    noInv,
		Status:       "open",
		TotalPrice:   float64(purchase.TotalPrice),
		TotalTebu:    float64(purchase.TotalTebu),
		TotalFarmer:  purchase.TotalFarmer,
		Note:         purchase.Note,
		PurchaseDate: purchase.PurchaseDate,
	}

	purchaseDetailModel := []purchaseigmmodel.PurchaseIgmDetail{}
	for _, value := range purchase.Items {
		totalHargaPurchase += float64(value.TotalKg) * float64(value.HargaKg)
		purchaseDetailModel = append(purchaseDetailModel, purchaseigmmodel.PurchaseIgmDetail{
			IDUser:       int(value.IDUser),
			IDUserLand:   int(value.IDUserLand),
			IDProduction: int(value.IDProduction),
			TotalKg:      float64(value.TotalKg),
			HargaKg:      float64(value.HargaKg),
			Subtotal:     float64(value.TotalKg) * float64(value.HargaKg),
		})
	}

	purchaseModel.TotalPrice = float64(totalHargaPurchase)
	err = t.PurchaseIgmRepository.Save(purchaseModel, purchaseDetailModel)
	if err != nil {
		return err
	}

	return nil
}

func (t *PurchaseIgmServiceImpl) Delete(purchaseId int) error {
	err := t.PurchaseIgmRepository.Delete(purchaseId)
	if err != nil {
		return err
	}
	return nil
}

func (t *PurchaseIgmServiceImpl) FindAll(purchase *purchaseigmmodel.PurchaseIgm) ([]response.PurchaseIgmResponse, error) {
	result, err := t.PurchaseIgmRepository.GetAllByQuery(*purchase)

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

	var purchases []response.PurchaseIgmResponse
	for _, value := range result {
		purchases = append(purchases, t.formattedResponse(value))
	}

	return purchases, nil
}

func (t *PurchaseIgmServiceImpl) FindById(purchaseId int) (response.PurchaseIgmResponse, error) {
	purchaseData, err := t.PurchaseIgmRepository.GetOneByQuery(purchaseigmmodel.PurchaseIgm{ID: uint(purchaseId)})
	if err != nil {
		return response.PurchaseIgmResponse{}, err
	}

	// get product detail
	formatResponse := t.formattedResponse(purchaseData)
	return formatResponse, nil
}

func (t *PurchaseIgmServiceImpl) formattedResponse(value purchaseigmmodel.PurchaseIgm) response.PurchaseIgmResponse {
	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return response.PurchaseIgmResponse{}
	}
	listUser := map[int]response.UserResponse{}
	for _, value := range users {
		listUser[int(value.ID)] = value
	}

	// get user land service
	userLands, err := t.UserRemoteRepository.GetLands()
	if err != nil {
		return response.PurchaseIgmResponse{}
	}
	listUserLand := map[int]userresponse.UserLandRowResponse{}
	for _, value := range userLands.Data {
		listUserLand[int(value.ID)] = value
	}

	// fetch purchaseIgm
	purchaseIgmDetailDb, err := t.PurchaseIgmDetailRepository.GetAllByQuery(purchaseigmmodel.PurchaseIgmDetail{IDPurchaseIgm: int(value.ID)})
	if err != nil {
		purchaseIgmDetailDb = []purchaseigmmodel.PurchaseIgmDetail{}
	}
	var purchaseIgmDetail []response.PurchaseIgmItemResponse
	for _, item := range purchaseIgmDetailDb {

		land := listUserLand[item.IDUserLand]

		purchaseIgmDetail = append(purchaseIgmDetail, response.PurchaseIgmItemResponse{
			IDUser:       item.IDUser,
			IDUserLand:   item.IDUserLand,
			IDProduction: item.IDProduction,
			TotalKg:      item.TotalKg,
			HargaKg:      item.HargaKg,
			Subtotal:     item.Subtotal,
			UserDetail: response.PurchaseIgmUserDetailResponse{
				UserResponse: listUser[int(item.IDUser)],
			},
			UserLandDetail: response.PurchaseIgmUserLandResponse{
				LandName:    land.LandName,
				LandAddress: land.LandAddress,
				LandArea:    land.LandArea,
				TotalObat:   land.TotalObat,
				TotalPupuk:  land.TotalPupuk,
				TotalBibit:  land.TotalBibit,
				TotalTebu:   land.TotalTebu,
			},
		})
	}

	return response.PurchaseIgmResponse{
		ID:           value.ID,
		NoInvoice:    value.NoInvoice,
		PurchaseDate: value.PurchaseDate,
		Note:         value.Note,
		TotalTebu:    value.TotalTebu,
		TotalPrice:   value.TotalPrice,
		TotalFarmer:  value.TotalFarmer,
		Status:       value.Status,
		Items:        purchaseIgmDetail,
	}
}
