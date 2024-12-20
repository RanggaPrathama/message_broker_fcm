package main

import (
	"fmt"

	"github.com/RanggaPrathama/message_broker_fcm/lib"
	//"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/domain/repository"
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/RanggaPrathama/message_broker_fcm/routes"
	"github.com/RanggaPrathama/message_broker_fcm/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	lib.ConnectionPostgree()

	// Migrate the schema
	//lib.Database.AutoMigrate(&models.Message{},&models.User{}, &models.DeviceUser{})


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userRepo := repository.NewUserRepository(lib.Database)
	userService := service.NewUserService(userRepo)
	handler := handler.NewUserHandler(userService)

	routes.UserRoute(app, handler)


	app.Listen(fmt.Sprintf(":%s", lib.LoadEnv("APP_PORT")))
}