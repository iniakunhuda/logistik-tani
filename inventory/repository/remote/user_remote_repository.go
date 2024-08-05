package remote

import (
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type UserRemoteRepository interface {
	Profile() (response.UserResponse, error)
}
