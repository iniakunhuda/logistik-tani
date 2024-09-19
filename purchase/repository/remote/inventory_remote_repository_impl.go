package remote

import (
	"errors"

	"github.com/imroc/req/v3"
	"github.com/iniakunhuda/logistik-tani/purchase/response"
	inventoryresponse "github.com/iniakunhuda/logistik-tani/purchase/response/inventory_response"
	panenresponse "github.com/iniakunhuda/logistik-tani/purchase/response/panen_response"
	"github.com/iniakunhuda/logistik-tani/purchase/util"
)

type InventoryRemoteRepositoryImpl struct {
	baseUrl     string
	bearerToken string
}

func NewInventoryRemoteRepositoryImpl(bearerToken string) InventoryRemoteRepository {
	return InventoryRemoteRepositoryImpl{
		baseUrl:     util.GetEnv("INVENTORY_SERVICE_BASE_URL", "http://localhost:4001/api"),
		bearerToken: bearerToken,
	}
}

func (t InventoryRemoteRepositoryImpl) GetAll() ([]response.ProductResponse, error) {
	var inventoryResponse inventoryresponse.InventoryListResponse
	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&inventoryResponse).
		Get(t.baseUrl + "/inventory/all")

	if err != nil {
		return nil, err
	}

	if resp.IsErrorState() {
		return nil, err
	}

	inventoryDetail := inventoryResponse.Data
	return inventoryDetail, nil
}

func (t InventoryRemoteRepositoryImpl) GetDetail(id string) (response.ProductResponse, error) {
	var inventoryResponse inventoryresponse.InventoryDetailResponse
	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&inventoryResponse).
		Get(t.baseUrl + "/inventory/all/detail/" + id)

	if err != nil {
		return response.ProductResponse{}, err
	}

	if resp.IsErrorState() {
		return response.ProductResponse{}, err
	}

	inventoryDetail := inventoryResponse.Data
	return inventoryDetail, nil
}

func (t InventoryRemoteRepositoryImpl) UpdateReduceStok(id string, stok string) error {
	var response inventoryresponse.InventoryDefaultResponse

	url := t.baseUrl + "/inventory/all/update_reduce_stock/" + id
	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetBody(map[string]interface{}{
			"id_produk":    id,
			"stok_terbaru": stok,
			"description":  "sales",
		}).
		SetSuccessResult(&response).
		SetErrorResult(&response).
		Put(url)

	if err != nil {
		return err
	}

	if resp.IsErrorState() {
		return errors.New(response.Message)
	}

	return nil
}

func (t InventoryRemoteRepositoryImpl) UpdateIncreaseStok(id string, stok string) error {
	var response inventoryresponse.InventoryDefaultResponse

	url := t.baseUrl + "/inventory/all/update_increase_stock/" + id
	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetBody(map[string]interface{}{
			"id_produk":    id,
			"stok_terbaru": stok,
			"description":  "purchase",
		}).
		SetSuccessResult(&response).
		SetErrorResult(&response).
		Put(url)

	if err != nil {
		return err
	}

	if resp.IsErrorState() {
		return errors.New(response.Message)
	}

	return nil
}

func (t InventoryRemoteRepositoryImpl) AutoCreateProdukPetani(produk response.ProductResponse, qty uint, idPembeli uint) error {
	var response inventoryresponse.InventoryDefaultResponse

	url := t.baseUrl + "/inventory/petani"

	// remove produk.ID
	produk.ID = 0
	produk.Stock = qty
	produk.IDUser = idPembeli

	// jcart, _ := json.Marshal(produk)
	// fmt.Println(string(jcart))
	// return errors.New("test")

	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetBody(produk).
		SetSuccessResult(&response).
		SetErrorResult(&response).
		Post(url)

	if err != nil {
		return err
	}

	if resp.IsErrorState() {
		return errors.New(response.Message)
	}

	return nil
}

func (t InventoryRemoteRepositoryImpl) GetPanenAll() ([]panenresponse.ProductionRowResponse, error) {
	var res panenresponse.ProductionListResponse
	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&res).
		Get(t.baseUrl + "/panen/all")

	if err != nil {
		return nil, err
	}

	if resp.IsErrorState() {
		return nil, err
	}

	inventoryDetail := res.Data
	return inventoryDetail, nil
}

func (t InventoryRemoteRepositoryImpl) GetPanenDetail(id string) (panenresponse.ProductionRowResponse, error) {
	var res panenresponse.ProductionDetailResponse
	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&res).
		Get(t.baseUrl + "/panen/all/detail/" + id)

	if err != nil {
		return panenresponse.ProductionRowResponse{}, err
	}

	if resp.IsErrorState() {
		return panenresponse.ProductionRowResponse{}, err
	}

	inventoryDetail := res.Data
	return inventoryDetail, nil
}
