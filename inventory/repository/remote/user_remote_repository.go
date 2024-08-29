package remote

import (
	"github.com/iniakunhuda/logistik-tani/inventory/response"
	userresponse "github.com/iniakunhuda/logistik-tani/inventory/response/user_response"
)

type UserRemoteRepository interface {
	Profile() (response.UserResponse, error)
	GetAll() ([]response.UserResponse, error)
	Find(id string) (response.UserResponse, error)
	GetLandByUserId(userId string, landId string) (userresponse.UserLandResponse, error)
}
