package remote

import (
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type UserRemoteRepositoryImpl struct {
}

func NewUserRemoteRepositoryImpl() UserRemoteRepositoryImpl {
	return UserRemoteRepositoryImpl{}
}

func (t UserRemoteRepositoryImpl) Profile() (response.UserResponse, error) {
	var userResponse response.UserResponse
	_, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&userResponse).
		Get("http://localhost:4000/api/users/profile")

	if err != nil {
		fmt.Println(err)
		return response.UserResponse{}, err
	}

	return response.UserResponse{}, nil
}
