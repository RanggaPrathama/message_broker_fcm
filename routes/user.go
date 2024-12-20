package routes

import (
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, handler *handler.UserHandler){

	api := app.Group("/api")
	api.Get("/users", handler.FindAllUser)
	api.Post("/users/create", handler.CreateUser)
}