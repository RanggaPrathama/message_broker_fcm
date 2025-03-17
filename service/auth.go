package service

import (
	//"errors"
	"fmt"

	//"time"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	URepo "github.com/RanggaPrathama/message_broker_fcm/domain/repository/interfaces"
	"github.com/RanggaPrathama/message_broker_fcm/utils"
	//"gorm.io/gorm"
)

type AuthService struct {
	userRepo URepo.User
	deviceRepo URepo.DeviceRepositoryInterface
}

func NewAuthService(repo URepo.User, deviceRepo URepo.DeviceRepositoryInterface) *AuthService {
	return &AuthService{
		userRepo: repo,
		deviceRepo: deviceRepo,
	}
}

func (auth *AuthService) Login(userRequest models.UserLoginRequest) (models.UserLoginResponse, error){

	var userResponse models.UserLoginResponse

	if userRequest.EMAIL == "" {
		return userResponse, fmt.Errorf("email is required")
	}

	usersRepo, err := auth.userRepo.FindUserByEmail(userRequest.EMAIL)

	if err != nil {
		return userResponse, fmt.Errorf("email is not registered")
	}

	
	if !utils.ComparePassword(usersRepo.PASSWORD, userRequest.PASSWORD){
		return userResponse, fmt.Errorf("password is incorrect")
	}

	fmt.Println("User Request" , userRequest)

	// if userRequest.DEVICE_IDPHONE == "" {
	// 	return userResponse, fmt.Errorf("device id is required")
	// }

	//deviceResponse, err := auth.deviceRepo.GetDeviceUserByActive(usersRepo.ID_USER)
	


	// if  deviceResponse.IS_ACTIVE && userRequest.DEVICE_IDPHONE != deviceResponse.DEVICE_ID_PHONE {
	// 	fmt.Println("Device login another device")
	// 	return userResponse, fmt.Errorf("sorry, you have logged in on another device")
	// }


	// if err != nil && errors.Is(err, gorm.ErrRecordNotFound){
	// 	fmt.Println("Device Not Found")
		
	// 	device := models.DeviceUser{
	// 	   USER_ID_USER: usersRepo.ID_USER,
	// 	   DEVICE_IP_ADDRESS: userRequest.DEVICE_IP_ADDRESS,
	// 	   DEVICE_ID_PHONE: userRequest.DEVICE_IDPHONE,
	// 	   PLATFORM: userRequest.PLATFORM,
	// 	   DEVICE_TOKEN: userRequest.DEVICE_TOKEN,
	// 	   CREATED_AT: func() *time.Time { t := time.Now(); return &t }(),
	// 	   IS_ACTIVE: true,
	// 	}

	// 	err = auth.deviceRepo.CreateDevice(device)

	// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return userResponse, err
	// 	}
	// }

	

	// if deviceResponse.DEVICE_TOKEN != userRequest.DEVICE_TOKEN{
	// 	if err := auth.deviceRepo.UpdateDeviceTokenFcm(u)
	// }
	

	token, err := utils.GenerateJwtToken(usersRepo.ID_USER, usersRepo.EMAIL)

	if err != nil {
		return userResponse, err
	}

	if err := auth.userRepo.UpdateLastLogin(usersRepo.ID_USER); err != nil {
		return userResponse, fmt.Errorf("failed to update last login: %w", err)
	}

	userResponse = models.UserLoginResponse{
		ID_USER: usersRepo.ID_USER,
		USERNAME: usersRepo.USERNAME,
		EMAIL: usersRepo.EMAIL,
		TOKEN: token,
		CREATED_AT: usersRepo.CREATED_AT,
		UPDATED_AT: usersRepo.UPDATED_AT,
	}
	

	return userResponse, nil

}


func (auth *AuthService) HandlerLoginCallback(userRequest models.GoogleLoginRequest) (models.UserLoginResponse, error){

	
    
  return models.UserLoginResponse{}, nil
}
