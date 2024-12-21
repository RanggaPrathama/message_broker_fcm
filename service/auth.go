package service

import (
	"fmt"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	URepo "github.com/RanggaPrathama/message_broker_fcm/domain/repository/interfaces"
	"github.com/RanggaPrathama/message_broker_fcm/utils"
)

type AuthService struct {
	userRepo URepo.User
}

func NewAuthService(repo URepo.User) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (auth *AuthService) Login(user models.UserLoginResponse) (models.UserLoginResponse, error){

	var userLogin models.UserLoginResponse

	if user.EMAIL == "" {
		return user, fmt.Errorf("email is required")
	}

	usersRepo, err := auth.userRepo.FindUserByEmail(user.EMAIL)

	if err != nil {
		return user, err
	}

	print(usersRepo.PASSWORD)
	if !utils.ComparePassword(usersRepo.PASSWORD, user.PASSWORD){
		return user, fmt.Errorf("password is incorrect")
	}

	token, err := utils.GenerateJwtToken(usersRepo.ID_USER, usersRepo.EMAIL)

	if err != nil {
		return user, err
	}

	userLogin = models.UserLoginResponse{
		ID_USER: usersRepo.ID_USER,
		USERNAME: usersRepo.USERNAME,
		EMAIL: usersRepo.EMAIL,
		TOKEN: token,
		CREATED_AT: usersRepo.CREATED_AT,
		UPDATED_AT: usersRepo.UPDATED_AT,
	}
	

	return userLogin, nil

}