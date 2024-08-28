package service

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/repository"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type ProductServiceImpl struct {
	ProductRepository      repository.ProductRepository
	ProductOwnerRepository repository.ProductOwnerRepository
	Validate               *validator.Validate
}

func NewProductServiceImpl(productRepo repository.ProductRepository, productOwnerRepo repository.ProductOwnerRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository:      productRepo,
		ProductOwnerRepository: productOwnerRepo,
		Validate:               validate,
	}
}

func (t *ProductServiceImpl) Create(request request.CreateProdukRequest) error {
	// check if product name exists, if exists fetch the product
	productDb, _ := t.ProductRepository.FindByName(request.Name)

	// if productdb nil, create new product
	if productDb == nil {
		newProductDb := model.Product{
			Name:        request.Name,
			Description: request.Description,
			PriceBuy:    request.PriceBuy,
			PriceSell:   request.PriceSell,
			Category:    request.Category,
		}
		err := t.ProductRepository.Save(newProductDb)
		if err != nil {
			return err
		}

		// fetch again
		productDb, err = t.ProductRepository.FindByName(request.Name)
		if err != nil {
			return err
		}
	}

	// create product owner
	newProductOwner := model.ProductOwner{
		IDUser:      int(request.IDUser),
		IDProduct:   int(productDb.ID),
		Stock:       request.Stock,
		PriceBuy:    request.PriceBuy,
		PriceSell:   request.PriceSell,
		Description: request.Description,
	}
	err := t.ProductOwnerRepository.Save(newProductOwner)
	if err != nil {
		return err
	}

	return nil
}

// func (t *ProductServiceImpl) Delete(produkId int) error {
// 	err := t.ProductRepository.Delete(produkId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (t *ProductServiceImpl) FindAll(produk *model.Product, userId ...string) ([]response.ProductResponse, error) {

	productDb, err := t.ProductOwnerRepository.GetAllByProduk(*produk, userId...)

	if err != nil {
		return nil, err
	}

	var produks []response.ProductResponse
	for _, value := range productDb {
		produk := response.ProductResponse{
			ID:          value.ID,
			IDUser:      uint(value.IDUser),
			Name:        value.Product.Name,
			Description: value.Product.Description,
			PriceBuy:    uint(value.PriceBuy),
			PriceSell:   uint(value.PriceSell),
			Category:    value.Product.Category,
			Stock:       uint(value.Stock),
		}
		produks = append(produks, produk)
	}

	return produks, nil
}

func (t *ProductServiceImpl) FindById(produkOwnerId int, userId ...string) (response.ProductResponse, error) {
	productDb, err := t.ProductOwnerRepository.GetAllByProduk(model.Product{ID: uint(produkOwnerId)}, userId...)
	if err != nil {
		return response.ProductResponse{}, err
	}

	if len(productDb) == 0 {
		return response.ProductResponse{}, nil
	}

	value := productDb[0]
	formatResponse := response.ProductResponse{
		ID:          value.ID,
		IDUser:      uint(value.IDUser),
		Name:        value.Product.Name,
		Description: value.Product.Description,
		PriceBuy:    uint(value.PriceBuy),
		PriceSell:   uint(value.PriceSell),
		Category:    value.Product.Category,
		Stock:       uint(value.Stock),
	}
	return formatResponse, nil
}

func (t *ProductServiceImpl) Update(produkOwnerId int, produk request.UpdateProdukRequest) (response.ProductResponse, error) {
	produkData, err := t.ProductOwnerRepository.FindById(produkOwnerId)
	if err != nil {
		return response.ProductResponse{}, err
	}

	// update all field
	if produk.Description != "" {
		produkData.Description = produk.Description
	}

	if produk.PriceBuy != 0 {
		produkData.PriceBuy = int(produk.PriceBuy)
	}

	if produk.PriceSell != 0 {
		produkData.PriceSell = int(produk.PriceSell)
	}

	if produk.Stock != 0 {
		produkData.Stock = int(produk.Stock)
	}

	t.ProductOwnerRepository.Update(*produkData)

	// find
	productDb, err := t.ProductOwnerRepository.GetAllByProduk(model.Product{ID: uint(produkOwnerId)})
	if err != nil {
		return response.ProductResponse{}, err
	}

	if len(productDb) == 0 {
		return response.ProductResponse{}, nil
	}

	value := productDb[0]
	formatResponse := response.ProductResponse{
		ID:          value.ID,
		IDUser:      uint(value.IDUser),
		Name:        value.Product.Name,
		Description: value.Product.Description,
		PriceBuy:    uint(value.PriceBuy),
		PriceSell:   uint(value.PriceSell),
		Category:    value.Product.Category,
		Stock:       uint(value.Stock),
	}

	return formatResponse, nil
}

func (t *ProductServiceImpl) UpdateReduceStock(productOwnerId int, stokTerbaru int) error {
	produkData, err := t.ProductOwnerRepository.FindById(productOwnerId)
	if err != nil {
		return err
	}

	if stokTerbaru > int(produkData.Stock) {
		return errors.New("Error! stok tidak mencukupi")
	}

	produkData.Stock = produkData.Stock - stokTerbaru
	t.ProductOwnerRepository.Update(*produkData)

	return nil
}

func (t *ProductServiceImpl) UpdateIncreaseStock(productOwnerId int, stokTerbaru int) error {
	produkData, err := t.ProductOwnerRepository.FindById(productOwnerId)
	if err != nil {
		return err
	}

	produkData.Stock = produkData.Stock + stokTerbaru
	t.ProductOwnerRepository.Update(*produkData)

	return nil
}

func (t *ProductServiceImpl) AutoCreateProductPetani(request request.CreateProdukRequest) error {
	// check if product name exists, if exists fetch the product
	productDb, _ := t.ProductRepository.FindByName(request.Name)

	fmt.Println("EAAAA", request.ID)

	// if productdb nil, create new product
	if productDb == nil {
		newProductDb := model.Product{
			Name:        request.Name,
			Description: request.Description,
			PriceBuy:    request.PriceBuy,
			PriceSell:   request.PriceSell,
			Category:    request.Category,
		}
		err := t.ProductRepository.Save(newProductDb)
		if err != nil {
			return err
		}

		// fetch again
		productDb, err = t.ProductRepository.FindByName(request.Name)
		if err != nil {
			return err
		}
	}

	// create product owner
	productOwnerDb, err := t.ProductOwnerRepository.GetOneByQuery(model.ProductOwner{IDProduct: int(productDb.ID), IDUser: int(request.IDUser)})
	if err != nil {
		return errors.New("Product Owner already exists")
	}

	if productOwnerDb != (model.ProductOwner{}) {
		productOwnerDb.Stock = productOwnerDb.Stock + request.Stock
		t.ProductOwnerRepository.Update(productOwnerDb)
	} else {
		newProductOwner := model.ProductOwner{
			IDUser:      int(request.IDUser),
			IDProduct:   int(productDb.ID),
			Stock:       request.Stock,
			PriceBuy:    request.PriceBuy,
			PriceSell:   request.PriceSell,
			Description: request.Description,
		}
		err := t.ProductOwnerRepository.Save(newProductOwner)
		if err != nil {
			return err
		}
	}

	return nil
}
