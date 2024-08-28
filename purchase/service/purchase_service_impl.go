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

type InventoryServiceImpl struct {
	TokenAuth                 string
	PurchaseRepository        repository.PurchaseRepository
	UserRemoteRepository      remote.UserRemoteRepository
	InventoryRemoteRepository remote.InventoryRemoteRepository
	Validate                  *validator.Validate
}

func NewInventoryServiceImpl(purchasesRepository repository.PurchaseRepository, validate *validator.Validate) PurchaseService {
	return &InventoryServiceImpl{
		PurchaseRepository:        purchasesRepository,
		UserRemoteRepository:      remote.NewUserRemoteRepositoryImpl(),
		InventoryRemoteRepository: remote.NewInventoryRemoteRepositoryImpl(""), // TODO: set bearer token
		Validate:                  validate,
	}
}

func (t *InventoryServiceImpl) GenerateNoInvoice() (string, error) {
	purchases, err := t.PurchaseRepository.FindLastRow()
	if err != nil {
		return "", err
	}

	lastInv := 0
	if purchases != nil {
		parts := strings.Split(purchases.NoInvoice, "-")
		if len(parts) == 2 {
			lastInv, err = strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
		}
	}

	noInv := ""
	if lastInv+1 < 10 {
		noInv = fmt.Sprintf("PURCHASE-000%d", lastInv+1)
	} else if lastInv+1 < 100 {
		noInv = fmt.Sprintf("PURCHASE-00%d", lastInv+1)
	} else if lastInv+1 < 1000 {
		noInv = fmt.Sprintf("PURCHASE-0%d", lastInv+1)
	}

	return noInv, nil
}

func (t *InventoryServiceImpl) Create(purchase request.CreatePurchaseRequest) error {

	noInv, err := t.GenerateNoInvoice()
	if err != nil {
		return err
	}
	totalHargaPurchase := 0
	purchaseModel := model.Purchase{
		NoInvoice:     noInv,
		IDPembeli:     purchase.IDPembeli,
		IDPenjual:     purchase.IDPenjual,
		NamaPenjual:   purchase.NamaPenjual,
		AlamatPenjual: purchase.AlamatPenjual,
		TelpPenjual:   purchase.TelpPenjual,
		Tanggal:       purchase.Tanggal,
		Status:        "open",
		TotalHarga:    0,
	}

	_, err = t.UserRemoteRepository.Find(strconv.Itoa(int(purchase.IDPembeli)))
	if err != nil {
		return err
	}

	purchaseDetailModel := []model.PurchaseDetail{}
	for _, value := range purchase.Produk {

		// Inventory Service: Check product stok
		inventoryDetail, err := t.InventoryRemoteRepository.GetDetail(strconv.Itoa(int(value.IDProduk)))
		if err != nil {
			return err
		}

		// Check produk.userId != inventory
		if inventoryDetail.IDUser != purchase.IDPembeli {
			return errors.New("Error! Produk tidak ditemukan")
		}

		purchaseDetailModel = append(purchaseDetailModel, model.PurchaseDetail{
			IDProduk:   value.IDProduk,
			Jenis:      value.Jenis,
			Harga:      int(value.Harga),
			Qty:        int(value.Qty),
			TotalHarga: int(value.Harga) * int(value.Qty),
		})

		// Inventory Service: Trigger stok update
		err = t.InventoryRemoteRepository.UpdateIncreaseStok(strconv.Itoa(int(value.IDProduk)), strconv.Itoa(int(value.Qty)))
		if err != nil {
			return err
		}

		totalHargaPurchase += int(value.Harga) * int(value.Qty)
	}

	purchaseModel.TotalHarga = totalHargaPurchase
	err = t.PurchaseRepository.Save(purchaseModel, purchaseDetailModel)
	if err != nil {
		return err
	}

	return nil
}

func (t *InventoryServiceImpl) Update(purchaseId int, userId int, purchases request.UpdatePurchaseRequest) error {
	// Validate the request
	purchasesData, err := t.PurchaseRepository.GetOneByQuery(model.Purchase{IDPembeli: uint(userId), ID: uint(purchaseId)})
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

func (t *InventoryServiceImpl) Delete(purchaseId int) error {
	err := t.PurchaseRepository.Delete(purchaseId)
	if err != nil {
		return err
	}
	return nil
}

func (t *InventoryServiceImpl) FindAll(purchase *model.Purchase) ([]response.PurchaseResponse, error) {
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

	var purchases []response.PurchaseResponse
	for _, value := range result {
		newPurchaseDetail := response.PurchaseResponse{
			Purchase:      value,
			IDPenjual:     value.IDPenjual,
			NamaPenjual:   value.NamaPenjual,
			AlamatPenjual: value.AlamatPenjual,
			TelpPenjual:   value.AlamatPenjual,
			// PenjualDetail: listUser[value.IDPenjual],
			PembeliDetail: listUser[value.IDPembeli],
		}
		purchases = append(purchases, newPurchaseDetail)
	}

	return purchases, nil
}

func (t *InventoryServiceImpl) FindById(purchaseId int, userId uint) (response.PurchaseResponse, error) {
	purchaseData, err := t.PurchaseRepository.GetOneByQuery(model.Purchase{IDPembeli: userId, ID: uint(purchaseId)})
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

	formatResponse := response.PurchaseResponse{
		Purchase:      purchaseData,
		IDPenjual:     purchaseData.IDPenjual,
		NamaPenjual:   purchaseData.NamaPenjual,
		AlamatPenjual: purchaseData.AlamatPenjual,
		TelpPenjual:   purchaseData.AlamatPenjual,
		// PenjualDetail: listUser[purchaseData.IDPenjual],
		PembeliDetail: listUser[purchaseData.IDPembeli],
	}
	return formatResponse, nil
}
