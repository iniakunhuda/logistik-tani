package response

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type ProductionResponse struct {
	model.Production
	UserDetail UserResponse `json:"user_detail"`
}
