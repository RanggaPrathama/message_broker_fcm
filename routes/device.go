package routes

import (
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/gofiber/fiber/v2"
)

func DeviceRoute(app *fiber.App, deviceHandler *handler.DeviceHandler) {
	api := app.Group("/api")
	api.Get("/devices", deviceHandler.FindAllDevice)
	api.Put("/devices/updateToken", deviceHandler.UpdateDeviceToken)
}