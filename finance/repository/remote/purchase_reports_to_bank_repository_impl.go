package remote

import (
	"fmt"

	"github.com/imroc/req/v3"
	purchasereportstobank "github.com/iniakunhuda/logistik-tani/finance/response/purchase_reports_to_bank"
	"github.com/iniakunhuda/logistik-tani/finance/util"
)

type PurchaseReportsToBankRepositoryImpl struct {
	baseUrl string
}

func NewPurchaseReportsToBankRepositoryImpl() PurchaseReportsToBankRepository {
	return PurchaseReportsToBankRepositoryImpl{
		baseUrl: util.GetEnv("PURCHASE_SERVICE_BASE_URL", "http://localhost:4003/api"),
	}
}

func (t PurchaseReportsToBankRepositoryImpl) GetAll() ([]purchasereportstobank.PurchaseReportsToBankResponse, error) {
	var responseBody purchasereportstobank.PurchaseReportsToBankApiListResponse
	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&responseBody).
		Get(t.baseUrl + "/report-to-bank")

	if err != nil {
		return []purchasereportstobank.PurchaseReportsToBankResponse{}, err
	}

	if resp.IsErrorState() {
		return []purchasereportstobank.PurchaseReportsToBankResponse{}, err
	}

	var userList []purchasereportstobank.PurchaseReportsToBankResponse
	userList = append(userList, responseBody.Data...)
	return userList, nil
}

func (t PurchaseReportsToBankRepositoryImpl) Find(id string) (purchasereportstobank.PurchaseReportsToBankResponse, error) {

	var responseBody purchasereportstobank.PurchaseReportsToBankApiDetailResponse
	_, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&responseBody).
		Get(t.baseUrl + "/report-to-bank/" + id)

	if err != nil {
		fmt.Println(err)
		return purchasereportstobank.PurchaseReportsToBankResponse{}, err
	}

	return responseBody.Data, err
}
