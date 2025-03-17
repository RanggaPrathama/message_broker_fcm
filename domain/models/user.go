package models

import "time"

type User struct {
	ID_USER    uint      `gorm:"primaryKey;autoIncrement" json:"id_user"`
	USERNAME   string    `gorm:"type:varchar(255);not null" json:"username"` 
	EMAIL      string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	GOOGLE_ID  *string    `gorm:"type:varchar(255);null;unique" json:"google_id"` 
	AVATAR_GOOGLE string  `gorm:"type:varchar(255);null" json:"avatar_google"`
	PASSWORD   string    `gorm:"type:varchar(255);null" json:"password"`
	LAST_LOGIN *time.Time `gorm:"type:timestamp with time zone" json:"last_login"`
	CREATED_AT *time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UPDATED_AT *time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
}


type UserRegistrationRequest struct{
	USERNAME   string    ` json:"username"`
	EMAIL      string    ` json:"email"`
	PASSWORD   string    ` json:"password"`
	CREATED_AT *time.Time ` json:"created_at"`
}

type UserRegistrationResponse struct{
	ID_USER    uint      ` json:"id_user"`
	USERNAME   string    ` json:"username"`
	EMAIL      string    ` json:"email"`
	PASSWORD   string    ` json:"password"`
	CREATED_AT *time.Time ` json:"created_at"`
}


type UserLoginRequest struct{
	EMAIL      string    ` json:"email"` 
	PASSWORD   string    ` json:"password"`
	// DEVICE_IDPHONE string `json:"device_idphone"`
	// DEVICE_IP_ADDRESS string `json:"device_ip_address"`
	// PLATFORM string `json:"platform"`
	// DEVICE_TOKEN string `json:"device_token"`
}


type UserLoginResponse struct{
	ID_USER    uint      ` json:"id_user"`
	USERNAME   string    ` json:"username"` 
	EMAIL      string    ` json:"email"` 
	PASSWORD   string    ` json:"password"`
	CREATED_AT *time.Time ` json:"created_at"`
	UPDATED_AT *time.Time ` json:"updated_at"`
	TOKEN    string    `json:"token"`
}


type GoogleLoginRequest struct{
	EMAIL 	string    ` json:"email"`
	NAME 	string    ` json:"name"`
	AVATAR 	string    ` json:"avatar"`
	
}