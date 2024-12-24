package models

import "time"

type DeviceUser struct {
	ID_DEVICE    uint      `gorm:"primaryKey;autoIncrement" json:"id_device"`
	USER_ID_USER uint      `gorm:"not null" json:"user_id_user"`
	PLATFORM     string    `gorm:"type:varchar(50);not null" json:"platform"`
	DEVICE_IP_ADDRESS    string    `gorm:"type:varchar(100);not null" json:"device_ip_address"` 
	DEVICE_ID_PHONE string	`gorm:"type:varchar(100);not null" json:"device_id_phone"`
	DEVICE_TOKEN string    `gorm:"type:varchar(255);not null" json:"device_token"` 
	CREATED_AT   *time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UPDATED_AT   *time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
	IS_ACTIVE    bool	  	`gorm:"default:true" json:"is_active"`
	User         User      `gorm:"foreignKey:USER_ID_USER;references:ID_USER" json:"user"`
}

type DeviceUserRequest struct{
	DEVICE_IDPHONE string `json:"device_idphone"`
	DEVICE_IP_ADDRESS string `json:"device_ip_address"`
	PLATFORM string `json:"platform"`
	DEVICE_TOKEN string `json:"device_token"`
}
