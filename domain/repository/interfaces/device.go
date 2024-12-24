package interfaces

import "github.com/RanggaPrathama/message_broker_fcm/domain/models"

type DeviceRepositoryInterface interface {
	GetDeviceAll() ([]models.DeviceUser, error)
	CreateDevice(device models.DeviceUser) error
	GetDeviceByToken(token string) (models.DeviceUser, error)
	GetDeviceUserByActive(userId uint) (models.DeviceUser, error)
	GetDeviceByIdPhone(userId uint, id string) (*models.DeviceUser, error)
	UpdateDevice(device models.DeviceUser) error
	UpdateDeviceTokenFcm(id string, token string) error
	Deactivedevice(userId uint, deviceId string) error
}