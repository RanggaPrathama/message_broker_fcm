package service

import (
	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	Urepository "github.com/RanggaPrathama/message_broker_fcm/domain/repository/interfaces"
	Uservice "github.com/RanggaPrathama/message_broker_fcm/service/interfaces"
) 

type UserService struct {
	UserRepository Urepository.User

}

func NewUserService(userRepository Urepository.User) Uservice.UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (u *UserService) FindAllUser() ([]models.User, error) {

	users, err := u.UserRepository.FindAllUser()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserService) FindUserById(id uint) (models.User, error) {
	
	users, err := u.UserRepository.FindUserById(id)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *UserService) CreateUser(user models.User) error {
	
	err := u.UserRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}



