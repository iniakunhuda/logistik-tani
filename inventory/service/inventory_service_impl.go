package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/repository"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type InventoryServiceImpl struct {
	InventoryRepository repository.InventoryRepository
	Validate            *validator.Validate
	UserId              string
}

func NewInventoryServiceImpl(inventoryRepository repository.InventoryRepository, validate *validator.Validate) InventoryService {
	return &InventoryServiceImpl{
		InventoryRepository: inventoryRepository,
		Validate:            validate,
		// UserRemoteRepository: userRemoteRepository,
	}
}

func (t *InventoryServiceImpl) SetUserId(userId string) {
	t.UserId = userId
}

func (t *InventoryServiceImpl) Create(produk request.CreateProdukRequest) error {
	produkModel := model.Produk{
		IDUser:     produk.IDUser,
		NamaProduk: produk.NamaProduk,
		Hpp:        produk.Hpp,
		HargaJual:  produk.HargaJual,
		StokAktif:  produk.StokAktif,
		Kategori:   produk.Kategori,
		Jenis:      produk.Jenis,
		Varietas:   produk.Varietas,
		Status:     produk.Status,
	}
	err := t.InventoryRepository.Save(produkModel)

	if err != nil {
		return err
	}

	return nil
}

func (t *InventoryServiceImpl) Delete(produkId int) error {
	err := t.InventoryRepository.Delete(produkId)
	if err != nil {
		return err
	}
	return nil
}

func (t *InventoryServiceImpl) FindAll(produk *model.Produk) ([]response.ProdukResponse, error) {
	result, err := t.InventoryRepository.GetAllByQuery(*produk)

	if err != nil {
		return nil, err
	}

	var produks []response.ProdukResponse
	for _, value := range result {
		produk := response.ProdukResponse{
			Produk: value,
		}
		produks = append(produks, produk)
	}

	return produks, nil
}

func (t *InventoryServiceImpl) FindById(produkId int) (response.ProdukResponse, error) {
	produkData, err := t.InventoryRepository.FindById(produkId)
	if err != nil {
		return response.ProdukResponse{}, err
	}

	formatResponse := response.ProdukResponse{
		Produk: *produkData,
	}
	return formatResponse, nil
}

func (t *InventoryServiceImpl) Update(produkId int, produk request.UpdateProdukRequest) error {
	produkData, err := t.InventoryRepository.FindById(produkId)
	if err != nil {
		return err
	}

	// update all field

	if produk.NamaProduk != "" {
		produkData.NamaProduk = produk.NamaProduk
	}
	if produk.Hpp != 0 {
		produkData.Hpp = produk.Hpp
	}
	if produk.HargaJual != 0 {
		produkData.HargaJual = produk.HargaJual
	}
	if produk.StokAktif != 0 {
		produkData.StokAktif = produk.StokAktif
	}
	if produk.Kategori != "" {
		produkData.Kategori = produk.Kategori
	}
	if produk.Jenis != "" {
		produkData.Jenis = produk.Jenis
	}
	if produk.Varietas != "" {
		produkData.Varietas = produk.Varietas
	}
	if produk.Status != "" {
		produkData.Status = produk.Status
	}

	t.InventoryRepository.Update(*produkData)

	return nil
}

func (t *InventoryServiceImpl) UpdateReduceStock(produkId int, stokTerbaru int) error {
	produkData, err := t.InventoryRepository.FindById(produkId)
	if err != nil {
		return err
	}

	if stokTerbaru > int(produkData.StokAktif) {
		return errors.New("stok tidak mencukupi")
	}

	produkData.StokAktif = produkData.StokAktif - uint(stokTerbaru)
	t.InventoryRepository.Update(*produkData)

	return nil
}
