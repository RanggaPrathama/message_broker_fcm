package handler

import (
	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/response"
	"github.com/RanggaPrathama/message_broker_fcm/service/interfaces"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService interfaces.AuthServiceInterface
}

func NewAuthHandler(authService interfaces.AuthServiceInterface) *AuthHandler{
	return &AuthHandler{
		authService: authService,
	}
}

func (auth *AuthHandler) Login(c *fiber.Ctx) error {
	
	var user models.UserLoginResponse

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GlobalResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
		})
	}

	user , err := auth.authService.Login(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to login",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.GlobalResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data: user,
	})
}