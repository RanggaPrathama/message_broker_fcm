package lib

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/gofiber/contrib/websocket"
)

// Client mewakili setiap pengguna yang terhubung ke WebSocket
type Client struct {
	UserID int64
	RoomID int64
	Conn   *websocket.Conn
	Mu     *sync.Mutex
}

// Hub mengelola koneksi WebSocket berdasarkan Room
type Hub struct {
	ROOM       map[int64]map[int64]*Client // ROOM[roomID][userID] = *Client
	Broadcast  chan models.MessageRequest
	Registered chan *Client
	Unregister chan *Client
	Mu         sync.Mutex
}

// Buat instance Hub
func NewHub() *Hub {
	return &Hub{
		ROOM:       make(map[int64]map[int64]*Client),
		Broadcast:  make(chan models.MessageRequest),
		Registered: make(chan *Client),
		Unregister: make(chan *Client),
	}
}


func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Registered:
			h.Mu.Lock()
			
			if _, roomExists := h.ROOM[client.RoomID]; !roomExists {
				h.ROOM[client.RoomID] = make(map[int64]*Client)
			}
			

			h.ROOM[client.RoomID][client.UserID] = client
			h.Mu.Unlock()
			log.Println("Client Connected - Room:", client.RoomID, ", User:", client.UserID)

		case client := <-h.Unregister:
			h.Mu.Lock()
			if room, roomExists := h.ROOM[client.RoomID]; roomExists {
				if _, userExists := room[client.UserID]; userExists {
					// Hapus user dari Room
					delete(h.ROOM[client.RoomID], client.UserID)
					if len(h.ROOM[client.RoomID]) == 0 {
						delete(h.ROOM, client.RoomID) 
					}
					client.Conn.Close()
					log.Println("Client Disconnected - Room:", client.RoomID, ", User:", client.UserID)
				}
			}
			h.Mu.Unlock()

		case message := <-h.Broadcast:
			h.Mu.Lock()
			if room, exists := h.ROOM[int64(message.ROOM_ID)]; exists {
				for _, client := range room {
					msg, _ := json.Marshal(message) 
					client.Mu.Lock()
					if err := client.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Println("Error sending message:", err)
						client.Conn.Close()
						delete(room, client.UserID)
					}
					client.Mu.Unlock()
				}
			}
			h.Mu.Unlock()
		}
	}
}
