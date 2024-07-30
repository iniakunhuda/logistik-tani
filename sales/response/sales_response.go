package response

import "github.com/iniakunhuda/logistik-tani/sales/model"

type SalesResponse struct {
	model.Sales
	Password *struct{} `json:"password,omitempty"`
}
