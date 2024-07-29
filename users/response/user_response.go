package response

import "github.com/iniakunhuda/logistik-tani/users/model"

type UserResponse struct {
	model.User
	Password *struct{} `json:"password,omitempty"`
}
