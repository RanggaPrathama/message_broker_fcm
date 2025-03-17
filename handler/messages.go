package handler

import (
	//"strconv"

	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/lib"
	"github.com/RanggaPrathama/message_broker_fcm/response"
	"github.com/gofiber/fiber/v2"
)




func CreateROOM(c *fiber.Ctx) error {
	var room models.RoomChat


	// Parsing request body ke struct room
	if err := c.BodyParser(&room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Failed to parse request",
			"data":    nil,
		})
	}

	log.Printf("Received Room Data: %+v", room)

	// Pastikan koneksi database sudah ada
	if lib.Database == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Database connection is not initialized",
			"data":    nil,
		})
	}

	// Perbaiki format query SQL
	query := "INSERT INTO room_chats (chat_name, chat_type, created_at) VALUES (?, ?, ?)"

	// Eksekusi query menggunakan Exec
	 err := lib.Database.Exec(query, room.CHAT_NAME, room.CHAT_TYPE, time.Now())
	if err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to create room",
			"data":    fmt.Sprintf("%v", err),
		})
	}

	// Debugging jumlah rows yang dimasukkan
	rowsAffected := lib.Database.RowsAffected
	log.Printf("Rows Inserted: %d", rowsAffected)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Success",
		"data":    "Room created",
	})
}


func GetRoom(c *fiber.Ctx) error {

	query := `SELECT * FROM room_chats`
	
	var rooms []models.RoomChat

	err := lib.Database.Raw(query).Scan(&rooms)

	if err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to get rooms",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    rooms,
	})
}


func CreateMessages(c *fiber.Ctx) error {
	
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to parse request",
			"data":    nil,
		})
	}

	if message.CONTENT == "" || message.USER_ID_USER == 0 || message.ROOM_ID_CHAT == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request",
			"data":    nil,
		})
	}


	// err := lib.Database.Raw("SELECT * FROM room_members WHERE room_id_chat = ? AND user_id_user = ?", message.ROOM_ID_CHAT, message.USER_ID_USER).Scan(&message)

	// if err.Error != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"status" : fiber.StatusBadRequest,
	// 		"message" : "room or user not found",
	// 		"data" : nil,
	// 	})
	// }


	if message.MESSAGE_TYPE < 0 || message.MESSAGE_TYPE > 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request",
			"data":    nil,
		})
	}

	query := `INSERT INTO messages ( room_id_chat, user_id_user, content, message_type, created_at ) VALUES (?, ?, ?, ?, ?)`

	err := lib.Database.Exec(query, message.ROOM_ID_CHAT, message.USER_ID_USER, message.CONTENT, message.MESSAGE_TYPE, time.Now())

	if err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to create message",
			"data":    fmt.Sprintf("%v", err.Error),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.GlobalResponse{
		Status:  fiber.StatusCreated,
		Message: "Success",
		Data:    "Message created",
	})

}

func UploadImageOrFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to parse file request",
			"data":    nil,
		})
	}
	room_id_chat := c.FormValue("room_id_chat")
	user_id_user := c.FormValue("user_id_user")
	message_type := c.FormValue("message_type")

	if  room_id_chat == "" || user_id_user == "" || message_type == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request",
			"data":    nil,
		})
	}

	roomIdtoUint, _ := strconv.ParseUint(room_id_chat, 10, 64)
	userIdtoUint, _ := strconv.ParseUint(user_id_user, 10, 64)
	messageTypetoInt, _ := strconv.Atoi(message_type)

	if messageTypetoInt != 1 && messageTypetoInt != 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid message type (1=image, 2=file)",
			"data":    nil,
		})
	}

	// create folder 
	year := strconv.Itoa(time.Now().Year())
	var foldertype string

	if messageTypetoInt == 1 {
		foldertype = "images"
	} else {
		foldertype = "files"
	}

	uploadPath := fmt.Sprintf("uploads/%s/user_%d/%s/", year, userIdtoUint, foldertype)

	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to create folder",
			"data":    nil,
		})
	}

	filepath := filepath.Join(uploadPath, file.Filename)

	if err := c.SaveFile(file, filepath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to save file",
			"data":    nil,
		})
	}

	query := `INSERT INTO messages ( room_id_chat, user_id_user, content, message_type, created_at ) VALUES (?, ?, ?, ?, ?)`
	res := lib.Database.Exec(query, roomIdtoUint, userIdtoUint, filepath, messageTypetoInt, time.Now())

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to create message",
			"data":    fmt.Sprintf("%v", err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.GlobalResponse{
		Status:  fiber.StatusCreated,
		Message: "Success",
		Data:    "Message created",
	})
}

func UpdateStatusRead(c *fiber.Ctx) error {
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to parse request",
			"data":    nil,
		})
	}

	if message.ID_MESSAGE == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request",
			"data":    nil,
		})
	}

	query := `UPDATE messages SET is_read = ? WHERE id_message = ?`

	err := lib.Database.Exec(query, true, message.ID_MESSAGE)
	if err.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to update message",
			"data":    fmt.Sprintf("%v", err.Error),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.GlobalResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    "Message updated",
	})
}
