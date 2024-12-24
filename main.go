package main

import (
	"fmt"

	"github.com/RanggaPrathama/message_broker_fcm/lib"
	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/domain/repository"
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/RanggaPrathama/message_broker_fcm/routes"
	"github.com/RanggaPrathama/message_broker_fcm/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	lib.ConnectionPostgree()

	lib.Database.AutoMigrate(&models.Message{},&models.User{}, &models.DeviceUser{})


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userRepo := repository.NewUserRepository(lib.Database)
	deviceRepo := repository.NewDeviceRepository(lib.Database)

	userService := service.NewUserService(userRepo)
	userhandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, deviceRepo)
	authHandler := handler.NewAuthHandler(authService)

	deviceService := service.NewDeviceService(deviceRepo, userRepo)
	deviceHandler := handler.NewDeviceHandler(deviceService)
	
	routes.UserRoute(app, userhandler)
	routes.AuthRoute(app, authHandler)
	routes.DeviceRoute(app, deviceHandler)

	app.Listen(fmt.Sprintf(":%s", lib.LoadEnv("APP_PORT")))
}