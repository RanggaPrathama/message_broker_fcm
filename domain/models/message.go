package models

import "time"

type Message struct {
	ID_MESSAGE               uint       `gorm:"primaryKey;autoIncrement" json:"id_message"`
	ROOM_ID_CHAT             uint       `gorm:"not null" json:"room_id_chat"`
	USER_ID_USER             uint       `gorm:"not null" json:"user_id_user"`
	CONTENT 	   			string     	`gorm:"type:text;not null" json:"content"`
	MESSAGE_TYPE            int8       	`gorm:"not null" json:"message_type"` // 0 = text, 1 = image, 2 = file
	IS_READ 				bool       `gorm:"default:false" json:"is_read"`
	CREATED_AT              *time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UPDATE_AT               *time.Time `gorm:"type:timestamp with time zone" json:"update_at"`
	User                    User       `gorm:"foreignKey:USER_ID_USER;references:ID_USER" json:"user"`
	RoomChat                RoomChat   `gorm:"foreignKey:ROOM_ID_CHAT;references:ID_ROOM_CHAT" json:"room_chat"`
}


type MessageRequest struct {
	USER_ID     uint   `json:user_id`
	ROOM_ID     uint   `json:room_id`
	CONTENT     string `json:"content"`
	MESSAGE_TYPE int8   `json:"message_type"`
	IS_READ	 bool   `json:"is_read"`
	CREATED_AT  *time.Time `json:"created_at"`
}