package repository

import (
	"time"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/domain/repository/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.User{
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) FindAllUser() ([]models.User, error){


	var users []models.User

	err := repo.db.Find(&users)

	if err.Error != nil {
		return nil, err.Error
	}

	return users, nil
}

func(repo *UserRepository) FindUserById( id uint) (models.User, error){


	var user models.User 

	if err := repo.db.Where("id_user = ?", id).First(&user).Error;  err != nil {
		return user , err
	}


	return user, nil

}

func(repo *UserRepository) FindUserByEmail(email string) (models.User, error){
	var user models.User 

	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}


	return user, nil
}


func(repo *UserRepository) CreateUser(user models.User)  error {

	userForm := models.User{
		USERNAME: user.USERNAME,
		EMAIL: user.EMAIL,
		PASSWORD: user.PASSWORD,
		CREATED_AT: func(t time.Time) *time.Time { return &t }(time.Now()),
	} 
	
	err := repo.db.Create(&userForm)

	if err.Error != nil {
		return err.Error
	}
	

	return nil
}

func(repo *UserRepository) UpdateLastLogin(id uint) error {

	err := repo.db.Model(models.User{}).Where("id_user = ?", id).Update("last_login", time.Now())

	if err.Error != nil {
		return err.Error
	}

	return nil
}