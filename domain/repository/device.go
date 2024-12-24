package repository

import (
	"fmt"
	"time"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/domain/repository/interfaces"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) interfaces.DeviceRepositoryInterface {
	return &DeviceRepository{
		db: db,
	}
}

func (repo *DeviceRepository) GetDeviceAll() ([]models.DeviceUser, error) {

	var devices []models.DeviceUser
	
	// result := repo.db.Table("device_users").Joins("JOIN users ON device_users.user_id = users.id").Select("device_users.id, device_users.token, device_users.user_id, users.email").Scan(&devices)
	result := repo.db.Model(&models.DeviceUser{}).Preload("User").Find(&devices)

	if result.Error != nil {
		return nil, result.Error
	}

	return devices, nil
}

func (repo *DeviceRepository) CreateDevice(device models.DeviceUser) error {
	
	//  var device models.DeviceUser

	result := repo.db.Create(&device)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *DeviceRepository) GetDeviceByToken(token string) (models.DeviceUser, error) {
	var device models.DeviceUser

	result := repo.db.Where("device_token = ?", token).First(&device)

	if result.Error != nil {
		return device, result.Error
	}

	return device, nil
}

func (repo *DeviceRepository) GetDeviceByIdPhone(user_id uint,id string) (*models.DeviceUser, error) {

	var device models.DeviceUser

	result := repo.db.Where("device_id_phone = ? AND user_id_user = ? ", id, user_id, true).First(&device)

	fmt.Println("DEVICE REPO ", device)

	if result.Error != nil {
		return &device, result.Error
	}

	return &device, nil
}

func(repo *DeviceRepository) GetDeviceUserByActive(userId uint) (models.DeviceUser, error) {
	var device models.DeviceUser

	result := repo.db.Where("is_active = ? AND user_id_user = ?", true, userId).First(&device)

	if result.Error != nil {
		return device, result.Error
	}

	return device, nil
}


func (repo *DeviceRepository) UpdateDevice(device models.DeviceUser) error {

	var deviceUpdate models.DeviceUser

	result := repo.db.Model(&deviceUpdate).Where("id_device = ?", device.ID_DEVICE).Updates(models.DeviceUser{
		PLATFORM: device.PLATFORM,
		DEVICE_IP_ADDRESS: device.DEVICE_IP_ADDRESS,
		DEVICE_TOKEN: device.DEVICE_TOKEN,
		IS_ACTIVE: device.IS_ACTIVE,
		UPDATED_AT: func(t time.Time) *time.Time { return &t }(time.Now()),
	})

	if result.Error != nil {
		return result.Error

	}

	return nil
}	


func(repo *DeviceRepository) UpdateDeviceTokenFcm(id string, token string) error {
	
	err := repo.db.Model(&models.DeviceUser{}).Where("device_id_phone = ?", id).Update("device_token", token)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (repo *DeviceRepository) Deactivedevice(userId uint, deviceId string) error {

	result := repo.db.Model(&models.DeviceUser{}).Where("user_id_user = ? AND id_device != ?", userId, deviceId).Update("is_active", false)

	if result.Error != nil {
		return result.Error
	}

	return nil
}