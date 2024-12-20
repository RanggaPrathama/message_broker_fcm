package interfaces

import (
	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
)

type User interface {
	FindAllUser() ([]models.User, error)
	FindUserById(id uint) (models.User, error)
	CreateUser(user models.User) error
}