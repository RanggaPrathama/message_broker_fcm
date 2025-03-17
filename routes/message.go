package routes

import (
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/gofiber/fiber/v2"
)

func MessageRoute(app *fiber.App){
	api := app.Group("/api")
	api.Post("/room", handler.CreateROOM)

	api.Post("/message", handler.CreateMessages)
	
	api.Get("/get-room", handler.GetRoom)
}