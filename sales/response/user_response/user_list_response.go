package userresponse

import "github.com/iniakunhuda/logistik-tani/sales/response"

type UserListResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    []response.UserResponse `json:"data"`
}
