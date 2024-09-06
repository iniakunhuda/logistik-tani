package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/purchase/model"
	"github.com/iniakunhuda/logistik-tani/purchase/repository"
	"github.com/iniakunhuda/logistik-tani/purchase/repository/remote"
	"github.com/iniakunhuda/logistik-tani/purchase/request"
	"github.com/iniakunhuda/logistik-tani/purchase/response"
)

type PurchaseServiceImpl struct {
	TokenAuth                 string
	PurchaseRepository        repository.PurchaseRepository
	UserRemoteRepository      remote.UserRemoteRepository
	InventoryRemoteRepository remote.InventoryRemoteRepository
	Validate                  *validator.Validate
}

func NewPurchaseServiceImpl(purchasesRepository repository.PurchaseRepository, validate *validator.Validate) PurchaseService {
	return &PurchaseServiceImpl{
		PurchaseRepository:        purchasesRepository,
		UserRemoteRepository:      remote.NewUserRemoteRepositoryImpl(),
		InventoryRemoteRepository: remote.NewInventoryRemoteRepositoryImpl(""), // TODO: set bearer token
		Validate:                  validate,
	}
}

func (t *PurchaseServiceImpl) GenerateNoInvoice() (string, error) {
	sales, err := t.PurchaseRepository.FindLastRow()

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
		noInv = fmt.Sprintf("SALES-000%d", lastInv+1)
	} else if lastInv+1 < 100 {
		noInv = fmt.Sprintf("SALES-00%d", lastInv+1)
	} else if lastInv+1 < 1000 {
		noInv = fmt.Sprintf("SALES-0%d", lastInv+1)
	}

	return noInv, nil
}

func (t *PurchaseServiceImpl) Create(purchase request.CreatePurchaseRequest) error {

	noInv, err := t.GenerateNoInvoice()
	if err != nil {
		return err
	}
	totalHargaPurchase := 0
	purchaseModel := model.Purchase{
		NoInvoice:     noInv,
		IDBuyer:       int(purchase.IDBuyer),
		IDSeller:      purchase.IDSeller,
		SellerName:    purchase.SellerName,
		SellerAddress: purchase.SellerAddress,
		SellerTelp:    purchase.SellerTelp,
		PurchaseDate:  purchase.PurchaseDate,
		Status:        "open",
		TotalPrice:    0,
	}

	// jcart, _ := json.Marshal(purchaseModel)
	// fmt.Println(string(jcart))

	_, err = t.UserRemoteRepository.Find(strconv.Itoa(int(purchase.IDBuyer)))
	if err != nil {
		return err
	}

	purchaseDetailModel := []model.PurchaseDetail{}
	for _, value := range purchase.Products {

		// Inventory Service: Check product stok
		inventoryDetail, err := t.InventoryRemoteRepository.GetDetail(strconv.Itoa(int(value.IDProduct)))
		if err != nil {
			return err
		}

		// Check produk.userId != inventory
		if inventoryDetail.IDUser != purchase.IDBuyer {
			return errors.New("error! Produk tidak ditemukan")
		}

		purchaseDetailModel = append(purchaseDetailModel, model.PurchaseDetail{
			IDProductOwner: int(value.IDProduct),
			Price:          float64(value.Price),
			Qty:            int(value.Qty),
			Subtotal:       float64(value.Price) * float64(value.Qty),
		})

		// Inventory Service: Trigger stok update
		err = t.InventoryRemoteRepository.UpdateIncreaseStok(strconv.Itoa(int(value.IDProduct)), strconv.Itoa(int(value.Qty)))
		if err != nil {
			return err
		}

		totalHargaPurchase += int(value.Price) * int(value.Qty)
	}

	purchaseModel.TotalPrice = float64(totalHargaPurchase)
	err = t.PurchaseRepository.Save(purchaseModel, purchaseDetailModel)
	if err != nil {
		return err
	}

	return nil
}

func (t *PurchaseServiceImpl) Update(purchaseId int, userId int, purchases request.UpdatePurchaseRequest) error {
	// Validate the request
	purchasesData, err := t.PurchaseRepository.GetOneByQuery(model.Purchase{IDBuyer: int(userId), ID: uint(purchaseId)})
	if err != nil {
		return err
	}

	if purchasesData.Status == "done" {
		return errors.New("pembelian sudah tertutup, tidak dapat diedit")
	}

	statusArr := []string{"open", "pending", "done", "cancel"}
	isValid := false
	for _, value := range statusArr {
		if value == purchases.Status {
			isValid = true
			break
		}
	}

	if !isValid {
		return errors.New("invalid status. Available status: (open, pending, done, cancel)")
	}

	// Update the purchases
	err = t.PurchaseRepository.Update(model.Purchase{ID: uint(purchaseId), Status: purchases.Status})
	if err != nil {
		return err
	}

	return nil
}

func (t *PurchaseServiceImpl) Delete(purchaseId int) error {
	err := t.PurchaseRepository.Delete(purchaseId)
	if err != nil {
		return err
	}
	return nil
}

func (t *PurchaseServiceImpl) FindAll(purchase *model.Purchase) ([]response.PurchaseResponse, error) {
	result, err := t.PurchaseRepository.GetAllByQuery(*purchase)

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

	var purchases []response.PurchaseResponse
	for _, value := range result {
		// get product detail
		var purchaseDetail = value.PurchaseDetail
		for i, purchaseDetailItem := range purchaseDetail {
			productDetail := listProducts[uint(purchaseDetailItem.IDProductOwner)]
			purchaseDetail[i].Name = productDetail.Name
			purchaseDetail[i].Description = productDetail.Description
		}
		value.PurchaseDetail = purchaseDetail

		newPurchaseDetail := response.PurchaseResponse{
			Purchase:      value,
			IDSeller:      uint(value.IDSeller),
			SellerName:    value.SellerName,
			SellerAddress: value.SellerAddress,
			SellerTelp:    value.SellerTelp,
			BuyerDetail:   listUser[uint(value.IDBuyer)],
		}
		purchases = append(purchases, newPurchaseDetail)
	}

	return purchases, nil
}

func (t *PurchaseServiceImpl) FindById(purchaseId int, userId uint) (response.PurchaseResponse, error) {
	purchaseData, err := t.PurchaseRepository.GetOneByQuery(model.Purchase{IDBuyer: int(userId), ID: uint(purchaseId)})
	if err != nil {
		return response.PurchaseResponse{}, err
	}

	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return response.PurchaseResponse{}, err
	}
	listUser := map[uint]response.UserResponse{}
	for _, value := range users {
		listUser[value.ID] = value
	}

	// get list of products
	products, err := t.InventoryRemoteRepository.GetAll()
	if err != nil {
		return response.PurchaseResponse{}, err
	}
	listProducts := map[uint]response.ProductResponse{}
	for _, value := range products {
		listProducts[value.ID] = value
	}

	// get product detail
	var purchaseDetail = purchaseData.PurchaseDetail
	for i, purchaseDetailItem := range purchaseDetail {
		productDetail := listProducts[uint(purchaseDetailItem.IDProductOwner)]
		purchaseDetail[i].Name = productDetail.Name
		purchaseDetail[i].Description = productDetail.Description
	}
	purchaseData.PurchaseDetail = purchaseDetail

	formatResponse := response.PurchaseResponse{
		Purchase:      purchaseData,
		IDSeller:      uint(purchaseData.IDSeller),
		SellerName:    purchaseData.SellerName,
		SellerAddress: purchaseData.SellerAddress,
		SellerTelp:    purchaseData.SellerTelp,
		BuyerDetail:   listUser[uint(purchaseData.IDBuyer)],
	}
	return formatResponse, nil
}
