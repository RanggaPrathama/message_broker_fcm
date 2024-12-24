package interfaces

import "github.com/RanggaPrathama/message_broker_fcm/domain/models"

type AuthServiceInterface interface {
	Login(models.UserLoginRequest) (models.UserLoginResponse, error)

}