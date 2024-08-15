package remote

import (
	"github.com/iniakunhuda/logistik-tani/purchase/response"
)

type UserRemoteRepository interface {
	GetAll() ([]response.UserResponse, error)
	Find(id string) (response.UserResponse, error)
}
