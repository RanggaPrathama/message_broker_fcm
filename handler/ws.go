package handler

import (
	//"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/lib"
	"github.com/gofiber/contrib/websocket"
)

// WebSocketHandler menangani koneksi WebSocket
func WebSocketHandler(hub *lib.Hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		
		roomIDStr := conn.Query("room_id")
		userIDStr := conn.Query("user_id")

		roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
		if err != nil {
			log.Println("Invalid RoomID")
			conn.Close()
			return
		}

		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			log.Println("Invalid UserID")
			conn.Close()
			return
		}

		// Buat Client baru
		client := &lib.Client{
			UserID: userID,
			RoomID: roomID,
			Conn:   conn,
			Mu:     &sync.Mutex{},
		}

		// Daftarkan client ke hub
		hub.Registered <- client
		defer func() {
			hub.Unregister <- client
			conn.Close()
		}()

		// Loop untuk membaca pesan dari WebSocket
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}

			// Log pesan masuk
			fmt.Printf("Message from Room %d - User %d: %s\n", roomID, userID, string(message))

			// Broadcast pesan hanya ke room yang sesuai
			hub.Broadcast <- models.MessageRequest{
				// RoomID:  roomID,
				// UserID:  userID,
				// Content: string(message),
				USER_ID: uint(userID),
				ROOM_ID: uint(roomID),
				CONTENT: string(message),
				
			}
		}
	}
}
