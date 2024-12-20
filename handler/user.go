package handler

import (
	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/response"
	"github.com/RanggaPrathama/message_broker_fcm/service/interfaces"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService interfaces.UserServiceInterface
}

func NewUserHandler(userService interfaces.UserServiceInterface) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}


func (u *UserHandler) FindAllUser(c *fiber.Ctx) error {
	users, err := u.UserService.FindAllUser()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.GlobalResponse{
			Status:  fiber.StatusNotFound,
			Message: "Data not found",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.GlobalResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    users,
	})
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GlobalResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
		})
	}

	err := u.UserService.CreateUser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to create user",
			Data:    err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.GlobalResponse{
		Status:  fiber.StatusCreated,
		Message: "Success",
		Data: err,
	})
}
