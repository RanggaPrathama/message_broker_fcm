package models

import "time"

type RoomMember struct {
	ID_ROOM_MEMBER uint       `gorm:"primaryKey;autoIncrement" json:"id_room_member"`
	USER_ID_USER   uint       `gorm:"not null" json:"user_id_user"`
	ROOM_ID_CHAT   uint       `gorm:"not null" json:"room_id_chat"`
	User           User       `gorm:"foreignKey:USER_ID_USER;references:ID_USER" json:"user"`
	RoomChat       RoomChat   `gorm:"foreignKey:ROOM_ID_CHAT;references:ID_ROOM_CHAT" json:"room_chat"`
	CREATED_AT     *time.Time  `gorm:"type:timestamp with time zone" json:"created_at"`
	UPDATE_AT      *time.Time `gorm:"type:timestamp with time zone" json:"update_at"`
}