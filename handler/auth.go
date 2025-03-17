package handler

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"net/http/httptest"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/response"
	"github.com/RanggaPrathama/message_broker_fcm/service/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth/gothic"
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
	
	var user models.UserLoginRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.GlobalResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
		})
	}

	fmt.Println("USER PARSE", user)

	usersResponse , err := auth.authService.Login(user)

	// if err != nil && err.Error() == "sorry, you have logged in on another device"{
	// 	return c.Status(fiber.StatusBadRequest).JSON(response.GlobalResponse{
	// 		Status:  fiber.StatusBadRequest,
	// 		Message: "Failed to login",
	// 		Data:    err.Error(),
	// 	})
	// }

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
		Data: usersResponse,
	})
}





// Hangler Login Google
func HandlerLogin(w http.ResponseWriter, r *http.Request, provider string){
	// provider := chi.URLParam(r, "provider")
	// set context provider


	log.Println("PROVIDER", provider)

	// type contextKey string
	// var providerKey contextKey = "provider"

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
	 fmt.Println("error sini: ", gothUser)
	} else {
	 gothic.BeginAuthHandler(w, r)
	}
}


// Handler Callback Google

func (auth *AuthHandler) HandlerLoginCallback(c *fiber.Ctx) error {

	type contextKey string
	const providerKey contextKey = "provider"

	provider := c.Params("provider")

		ctx := context.WithValue(c.Context(), providerKey, provider)
	
		// Membuat http.Request dari Fiber request
		r, err := http.NewRequest(c.Method(), c.OriginalURL(), bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Gagal membuat request",
			})
		}
	
		// Menyalin header dari Fiber request ke http.Request
		for key, values := range c.GetReqHeaders() {
			for _, value := range values {
				r.Header.Add(key, value)
			}
		}
	
		r = r.WithContext(ctx)      
		w := httptest.NewRecorder() 
	
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}

		log.Print("USER", user)

		userRequest := models.GoogleLoginRequest{
			EMAIL: user.Email,
			NAME: user.Name,
			AVATAR: user.AvatarURL,
		}

     	 users , err := auth.authService.HandlerLoginCallback(userRequest)


		 userResponse := fiber.Map{
			"userRequest" : userRequest,
			"users" : users,
		 }

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.GlobalResponse{
				Status:  fiber.StatusBadRequest,
				Message: "Failed to login",
				Data:    err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(response.GlobalResponse{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data: userResponse,
		})
}