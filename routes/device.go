package routes

import (
	"github.com/RanggaPrathama/message_broker_fcm/handler"
	"github.com/RanggaPrathama/message_broker_fcm/middleware"
	"github.com/gofiber/fiber/v2"
)

func DeviceRoute(app *fiber.App, deviceHandler *handler.DeviceHandler) {
	api := app.Group("/api")
	devices := api.Group("/devices", middleware.VerifyJwtToken)
	devices.Get("/", deviceHandler.FindAllDevice)
	devices.Put("/updateToken", deviceHandler.UpdateDeviceToken)
	devices.Post("/cek-active", deviceHandler.CekDeviceActive)
}