package userresponse

import "github.com/iniakunhuda/logistik-tani/purchase/response"

type UserDefaultResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    response.UserResponse `json:"data"`
}
