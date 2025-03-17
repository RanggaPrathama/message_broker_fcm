package routes

import (
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/RanggaPrathama/message_broker_fcm/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, handler *handler.UserHandler) {

	api := app.Group("/api")
	
	api.Post("/registrasi", handler.CreateUser)
	
	users := api.Group("/users", middleware.VerifyJwtToken)
	users.Get("/", handler.FindAllUser)
	// users.Post("/create", handler.CreateUser)
}
