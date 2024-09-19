package remote

import (
	"github.com/iniakunhuda/logistik-tani/finance/response"
	userresponse "github.com/iniakunhuda/logistik-tani/finance/response/user_response"
)

type UserRemoteRepository interface {
	GetAll() ([]response.UserResponse, error)
	GetLands() (userresponse.UserLandResponse, error)
	Find(id string) (response.UserResponse, error)

	AddSaldo(id string, saldo int) error
}
