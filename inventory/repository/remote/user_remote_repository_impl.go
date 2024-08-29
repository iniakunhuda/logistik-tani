package remote

import (
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
	userresponse "github.com/iniakunhuda/logistik-tani/inventory/response/user_response"
	"github.com/iniakunhuda/logistik-tani/inventory/util"
)

type UserRemoteRepositoryImpl struct {
	baseUrl string
}

func NewUserRemoteRepositoryImpl() UserRemoteRepository {
	return UserRemoteRepositoryImpl{
		baseUrl: util.GetEnv("USER_SERVICE_BASE_URL", "http://localhost:4000/api"),
	}
}

func (t UserRemoteRepositoryImpl) Profile() (response.UserResponse, error) {
	var userResponse response.UserResponse
	_, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&userResponse).
		Get(t.baseUrl + "/users/profile")

	if err != nil {
		fmt.Println(err)
		return response.UserResponse{}, err
	}

	return response.UserResponse{}, nil
}

func (t UserRemoteRepositoryImpl) GetAll() ([]response.UserResponse, error) {
	var userResponse userresponse.UserListResponse
	resp, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&userResponse).
		Get(t.baseUrl + "/users")

	if err != nil {
		return []response.UserResponse{}, err
	}

	if resp.IsErrorState() {
		return []response.UserResponse{}, err
	}

	var userList []response.UserResponse
	userList = append(userList, userResponse.Data...)
	return userList, nil
}

func (t UserRemoteRepositoryImpl) Find(id string) (response.UserResponse, error) {

	var userResponse userresponse.UserDefaultResponse
	_, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&userResponse).
		Get(t.baseUrl + "/users/" + id)

	if err != nil {
		fmt.Println(err)
		return response.UserResponse{}, err
	}

	return userResponse.Data, err
}

func (t UserRemoteRepositoryImpl) GetLandByUserId(userId string, landId string) (userresponse.UserLandResponse, error) {
	var userResponse userresponse.UserLandResponse
	_, err := req.C().R().
		SetHeader("Accept", "application/json").
		SetBearerAuthToken("").
		EnableDump().
		SetSuccessResult(&userResponse).
		Get(t.baseUrl + "/lands?id_user=" + userId + "&id_land=" + landId)

	if err != nil {
		fmt.Println(err)
		return userResponse, err
	}

	return userResponse, nil
}
