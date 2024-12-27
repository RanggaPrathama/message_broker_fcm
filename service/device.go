package service

import (

	//"errors"
	"fmt"
	"time"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	DRepository "github.com/RanggaPrathama/message_broker_fcm/domain/repository/interfaces"
	DService "github.com/RanggaPrathama/message_broker_fcm/service/interfaces"
	//"gorm.io/gorm"
	
)

type DeviceService struct {
	deviceRepo DRepository.DeviceRepositoryInterface
	userRepo DRepository.User
}

func NewDeviceService(deviceRepo DRepository.DeviceRepositoryInterface, userRepo DRepository.User) DService.DeviceServiceInterface {
	return &DeviceService{
		deviceRepo: deviceRepo,
		userRepo: userRepo,
	}
}


func (service *DeviceService) FindAllDevice() ([]models.DeviceUser, error) {
	var devices []models.DeviceUser

	devices, err := service.deviceRepo.GetDeviceAll()

	if err != nil {
		return nil, err
	}

	return devices, nil

}

func (service *DeviceService) FindDeviceById(id uint) (models.DeviceUser, error) {
	var device models.DeviceUser

	//device, err := service.deviceRepo.GetDeviceByIdPhone(id)

	// if err != nil {
	// 	return device, err
	// }

	return device, nil	
}	

func (service *DeviceService) CreateDevice(device models.DeviceUser) error {

	err := service.deviceRepo.CreateDevice(device)

	if err != nil {
		return err
	}

	return nil
}

func (service *DeviceService) UpdateDevice(device models.DeviceUser) error {
	
	err := service.deviceRepo.UpdateDevice(device)

	if err != nil {
		return err
	}

	return nil
}


func (service *DeviceService) CreateOrUpdateDevice(device models.DeviceUser) error {

	//_, err := service.deviceRepo.GetDeviceByToken(device.DEVICE_TOKEN)



	// 
	return nil
}


func (service *DeviceService) UpdateDeviceTokenFcm(id string, token string) error {
	
	err := service.deviceRepo.UpdateDeviceTokenFcm(id, token)

	if err != nil {
		return err
	}

	return nil
}

func (service *DeviceService) CekDevice(userId uint, deviceRequest models.DeviceUserRequest) (models.DeviceUser, error) {

	device, err := service.deviceRepo.GetDeviceUserByActive(userId)
	
	// if err != nil {
	// 	return device, err
	// }

	print("Device ACTIVE", device.IS_ACTIVE)

	if deviceRequest.DEVICE_ID_PHONE == "" {
		return device, fmt.Errorf("device id phone is required")
	}

	if  err == nil && device.IS_ACTIVE &&  deviceRequest.DEVICE_ID_PHONE != device.DEVICE_ID_PHONE {
		return device, fmt.Errorf("sorry, you have logged in on another device")
	}

	//cek  device apakah tersimpan di database
	existingDevice, err := service.deviceRepo.GetDeviceByIdPhone(userId, deviceRequest.DEVICE_ID_PHONE)

	// if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
    //     return device, err
    // }

	fmt.Println("EXISTING DEVICE", existingDevice)
	if  err == nil && existingDevice.DEVICE_ID_PHONE != ""  {
		if err := service.deviceRepo.UpdateDevice(
			models.DeviceUser{
				DEVICE_ID_PHONE: deviceRequest.DEVICE_ID_PHONE,
				PLATFORM: deviceRequest.PLATFORM,
				DEVICE_IP_ADDRESS: deviceRequest.DEVICE_IP_ADDRESS,
				DEVICE_TOKEN: deviceRequest.DEVICE_TOKEN,
				IS_ACTIVE: true,
				UPDATED_AT: func() *time.Time { t := time.Now(); return &t}(),
			},
		); err != nil {
			return device, err
		}
	}else{
		device := models.DeviceUser{
			USER_ID_USER: userId,
			DEVICE_IP_ADDRESS: deviceRequest.DEVICE_IP_ADDRESS,
			DEVICE_ID_PHONE: deviceRequest.DEVICE_ID_PHONE,
			PLATFORM: deviceRequest.PLATFORM,
			DEVICE_TOKEN: deviceRequest.DEVICE_TOKEN,
			CREATED_AT: func() *time.Time { t := time.Now(); return &t }(),
			IS_ACTIVE: true,
		}

		err = service.deviceRepo.CreateDevice(device)

		if err != nil {
			return device, err
		}
	}

	// deactive device all agnore current device
	if err := service.deviceRepo.Deactivedevice(userId, deviceRequest.DEVICE_ID_PHONE); err != nil {
		return device, err
	}

	return device, nil
}