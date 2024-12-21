package routes

import (
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, authHandler *handler.AuthHandler){
	api := app.Group("/api")
	api.Post("/auth/login", authHandler.Login)
}