package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/repository"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type ProductServiceImpl struct {
	ProductRepository          repository.ProductRepository
	ProductOwnerRepository     repository.ProductOwnerRepository
	StockTransactionRepository repository.StockTransactionRepository
	Validate                   *validator.Validate
}

func NewProductServiceImpl(productRepo repository.ProductRepository, productOwnerRepo repository.ProductOwnerRepository, stockTransRepo repository.StockTransactionRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository:          productRepo,
		ProductOwnerRepository:     productOwnerRepo,
		StockTransactionRepository: stockTransRepo,
		Validate:                   validate,
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
		productDb, _ = t.ProductRepository.FindByName(request.Name)
		// if err != nil {
		// 	return err
		// }
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

	// add stock movement
	productOwnerDb, _ := t.ProductOwnerRepository.GetOneByQuery(model.ProductOwner{IDProduct: int(productDb.ID), IDUser: int(request.IDUser)})
	t.storeStockMovement(int(productOwnerDb.ID), int(request.IDUser), request.Stock, "init")

	return nil
}

func (t *ProductServiceImpl) FindAll(produk *model.Product, userId ...string) ([]response.ProductResponse, error) {

	productDb, err := t.ProductOwnerRepository.GetAllByProduk(*produk, userId...)

	if err != nil {
		return nil, err
	}

	var produks []response.ProductResponse
	for _, value := range productDb {

		produk := response.ProductResponse{
			ID:        value.ID,
			IDUser:    uint(value.IDUser),
			Name:      value.Product.Name,
			PriceBuy:  uint(value.PriceBuy),
			PriceSell: uint(value.PriceSell),
			Category:  value.Product.Category,
			Stock:     uint(value.Stock),
		}

		if value.Description != "" {
			produk.Description = value.Description
		} else {
			produk.Description = value.Product.Description
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
		ID:        value.ID,
		IDUser:    uint(value.IDUser),
		Name:      value.Product.Name,
		PriceBuy:  uint(value.PriceBuy),
		PriceSell: uint(value.PriceSell),
		Category:  value.Product.Category,
		Stock:     uint(value.Stock),
	}

	if value.Description != "" {
		formatResponse.Description = value.Description
	} else {
		formatResponse.Description = value.Product.Description
	}
	return formatResponse, nil
}

func (t *ProductServiceImpl) Update(produkOwnerId int, produk request.UpdateProdukRequest) (response.ProductResponse, error) {
	productOwnerDb, err := t.ProductOwnerRepository.FindById(produkOwnerId)
	if err != nil {
		return response.ProductResponse{}, err
	}
	jsonResponse, _ := json.Marshal(productOwnerDb)
	fmt.Println(string(jsonResponse))

	// update all field
	if produk.Name != "" {
		// check if product name exists, if exists fetch the product
		productDb, _ := t.ProductRepository.FindByName(produk.Name)
		if productDb == nil {
			// create product if not exists
			newProductDb := model.Product{
				Name:        produk.Name,
				Description: productOwnerDb.Product.Description,
				PriceBuy:    int(productOwnerDb.PriceBuy),
				PriceSell:   int(productOwnerDb.PriceSell),
				Category:    productOwnerDb.Product.Category,
			}
			err = t.ProductRepository.Save(newProductDb)
			if err != nil {
				return response.ProductResponse{}, err
			}

			// fetch again
			productDb, _ = t.ProductRepository.FindByName(produk.Name)
		}
		productOwnerDb.IDProduct = int(productDb.ID)

	}

	if produk.Description != "" {
		productOwnerDb.Description = produk.Description
	}

	if produk.PriceBuy != 0 {
		productOwnerDb.PriceBuy = int(produk.PriceBuy)
	}

	if produk.PriceSell != 0 {
		productOwnerDb.PriceSell = int(produk.PriceSell)
	}

	t.ProductOwnerRepository.Update(*productOwnerDb)

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
	}

	return formatResponse, nil
}

func (t *ProductServiceImpl) Delete(produkOwnerId int) error {
	produkData, err := t.ProductOwnerRepository.FindById(produkOwnerId)
	if err != nil {
		return err
	}

	// cek sales if exists return error
	_, err = t.StockTransactionRepository.GetOneByQuery(model.StockTransaction{IDProductOwner: produkOwnerId, Description: "sales"})
	if err == nil {
		return errors.New("error! Produk tidak bisa dihapus karena sudah ada transaksi penjualan/pembelian")
	}

	// cek purchase if exists return error
	_, err = t.StockTransactionRepository.GetOneByQuery(model.StockTransaction{IDProductOwner: produkOwnerId, Description: "purchase"})
	if err == nil {
		return errors.New("error! Produk tidak bisa dihapus karena sudah ada transaksi penjualan/pembelian")
	}

	// cek panen if exists return error
	_, err = t.StockTransactionRepository.GetOneByQuery(model.StockTransaction{IDProductOwner: produkOwnerId, Description: "panen"})
	if err == nil {
		return errors.New("error! Produk tidak bisa dihapus karena sudah ada transaksi penjualan/pembelian")
	}

	// delete stock movement
	err = t.StockTransactionRepository.DeleteByQuery(model.StockTransaction{IDProductOwner: produkOwnerId})
	if err != nil {
		return errors.New("error! Produk tidak bisa dihapus. Error: " + err.Error())
	}

	// delete product owner
	err = t.ProductOwnerRepository.Delete(int(produkData.ID))
	if err != nil {
		return errors.New("error! Produk tidak bisa dihapus. Error: " + err.Error())
	}

	return nil
}

func (t *ProductServiceImpl) UpdateReduceStock(productOwnerId int, stokTerbaru int, desc string) error {
	produkData, err := t.ProductOwnerRepository.FindById(productOwnerId)
	if err != nil {
		return err
	}

	if stokTerbaru > int(produkData.Stock) {
		return errors.New("error! stok tidak mencukupi")
	}

	produkData.Stock = produkData.Stock - stokTerbaru
	t.ProductOwnerRepository.Update(*produkData)

	// add stock movement
	t.storeStockMovement(productOwnerId, int(produkData.IDUser), stokTerbaru*-1, desc)

	return nil
}

func (t *ProductServiceImpl) UpdateIncreaseStock(productOwnerId int, stokTerbaru int, desc string) error {
	produkData, err := t.ProductOwnerRepository.FindById(productOwnerId)
	if err != nil {
		return err
	}

	produkData.Stock = produkData.Stock + stokTerbaru
	t.ProductOwnerRepository.Update(*produkData)

	// add stock movement
	t.storeStockMovement(productOwnerId, int(produkData.IDUser), stokTerbaru, desc)

	return nil
}

func (t *ProductServiceImpl) AutoCreateProductPetani(request request.CreateProdukRequest) error {
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
	productOwnerDb, _ := t.ProductOwnerRepository.GetOneByQuery(model.ProductOwner{IDProduct: int(productDb.ID), IDUser: int(request.IDUser)})
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

	// fetch again
	productOwnerDb, _ = t.ProductOwnerRepository.GetOneByQuery(model.ProductOwner{IDProduct: int(productDb.ID), IDUser: int(request.IDUser)})

	// add stock movement
	t.storeStockMovement(int(productOwnerDb.ID), int(request.IDUser), request.Stock, "purchase")

	return nil
}

func (t *ProductServiceImpl) storeStockMovement(idProductOwner int, idUser int, stock int, desc string) {
	stockTransaction := model.StockTransaction{
		IDProductOwner: idProductOwner,
		IDUser:         idUser,
		StockMovement:  stock,
		Date:           time.Now(),
		Description:    desc,
	}
	t.StockTransactionRepository.Save(stockTransaction)
}
