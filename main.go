package main

import (
	"fmt"

	"github.com/RanggaPrathama/message_broker_fcm/configs"
	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.ConnectionPostgree()

	// Migrate the schema
	configs.Database.AutoMigrate(&models.Message{},&models.User{}, &models.DeviceUser{})


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":%s", configs.LoadEnv("APP_PORT")))
}