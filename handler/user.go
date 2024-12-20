package handler

import (
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
