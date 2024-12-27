package middleware

import (
	"fmt"
	//"os"
	"strings"
	"time"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/lib"
	"github.com/RanggaPrathama/message_broker_fcm/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwtToken(c *fiber.Ctx) error {

	secretKey := lib.LoadEnv("SECRET_KEY")
	
	if secretKey == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(response.GlobalResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Secret Key is not defined",
	})
}
	authHeader := c.Get("Authorization")

	if authHeader == "" && !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(response.GlobalResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Unauthorized ",
			Data:    nil,
		})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenStr, &models.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secretKey), nil
	}  )
	
	

	claims, ok := token.Claims.(*models.JwtClaims)

	if !ok || !token.Valid || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.GlobalResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    nil,
		})
	}

	if claims.ExpiresAt.Time.Before(time.Now()){
		return c.Status(fiber.StatusUnauthorized).JSON(response.GlobalResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Token Expired",
			Data:    nil,
		})
	}

	c.Locals("id_user", claims.UserId)
	c.Locals("email", claims.Email)

	fmt.Println("USER ID", claims.UserId)


	return c.Next()
}