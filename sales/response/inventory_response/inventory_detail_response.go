package inventoryresponse

import "github.com/iniakunhuda/logistik-tani/sales/response"

type InventoryDetailResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    response.ProductResponse `json:"data"`
}
