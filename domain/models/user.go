package models

import "time"

type User struct {
	ID_USER    uint      `gorm:"primaryKey;autoIncrement" json:"id_user"`
	USERNAME   string    `gorm:"type:varchar(255);not null" json:"username"` 
	EMAIL      string    `gorm:"type:varchar(255);not null;unique" json:"email"` 
	PASSWORD   string    `gorm:"type:varchar(255);not null" json:"password"`
	CREATED_AT *time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UPDATED_AT *time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
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