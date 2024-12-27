package handler

import (
	"fmt"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/response"
	Dservice "github.com/RanggaPrathama/message_broker_fcm/service/interfaces"
	"github.com/gofiber/fiber/v2"
)

type DeviceHandler struct {
	DeviceService Dservice.DeviceServiceInterface
}

func NewDeviceHandler(deviceService Dservice.DeviceServiceInterface) *DeviceHandler {
	return &DeviceHandler{deviceService}
}

func (handler *DeviceHandler) FindAllDevice(c *fiber.Ctx) error {
	devices, err := handler.DeviceService.FindAllDevice()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to get devices",
			Data:    nil,
		})
	}


	return c.Status(fiber.StatusOK).JSON(response.GlobalResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    devices,
	})
}

func(handler *DeviceHandler) CreateDevice(c *fiber.Ctx) error {
	var device models.DeviceUser

	if err := c.BodyParser(&device); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to parse request",
			Data:    nil,
		})
	}

	err := handler.DeviceService.CreateDevice(device)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to create device",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.GlobalResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:   "Device created",
	})
}

func(handler *DeviceHandler) UpdateDeviceToken(c *fiber.Ctx) error {
	var device models.DeviceUser

	if err := c.BodyParser(&device); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to parse request",
			Data:    nil,
		})
	}

	err := handler.DeviceService.UpdateDeviceTokenFcm(device.DEVICE_ID_PHONE, device.DEVICE_TOKEN)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to update device",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.GlobalResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:   "Device updated",})

	}

	func (handler *DeviceHandler) CekDeviceActive(c *fiber.Ctx) error {

		var device models.DeviceUserRequest

		userId := c.Locals("id_user").(uint)
		fmt.Println("USER ID HANDLER", userId)

		if err := c.BodyParser(&device); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Failed to parse request",
				Data:    nil,
			})
		}

		deviceResponse, err := handler.DeviceService.CekDevice(userId,device)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
				Status:  fiber.StatusInternalServerError,
				Message: fmt.Sprintf("Failed to get device %v", err),
				Data:    err,
			})
		}


		return c.Status(fiber.StatusOK).JSON(response.GlobalResponse{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data:    deviceResponse,
		})

	}