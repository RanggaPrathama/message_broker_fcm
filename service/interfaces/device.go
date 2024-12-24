package interfaces

import "github.com/RanggaPrathama/message_broker_fcm/domain/models"

type DeviceServiceInterface interface {
	FindAllDevice() ([]models.DeviceUser, error)
	FindDeviceById(id uint) (models.DeviceUser, error)
	CreateOrUpdateDevice(device models.DeviceUser) error
	CreateDevice(device models.DeviceUser) error
	UpdateDevice(device models.DeviceUser) error
	UpdateDeviceTokenFcm(id string, token string) error
	CekDevice(userId uint, device models.DeviceUserRequest) (models.DeviceUser, error)
}