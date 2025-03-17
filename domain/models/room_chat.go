package models

import "time"


type RoomChat struct {
	ID_ROOM_CHAT uint `gorm:"primaryKey;autoIncrement" json:"id_room_chat"`
	CHAT_NAME *string `gorm:"type:varchar(255);column:chat_name;" json:"chat_name"`
	CHAT_TYPE bool `gorm:"column:chat_type;not null;default:0" json:"chat_type"`
	CREATED_AT *time.Time `gorm:"type:timestamp with time zone;column:created_at" json:"created_at"`
	UPDATE_AT *time.Time `gorm:"type:timestamp with time zone;column:update_at" json:"update_at"`
}