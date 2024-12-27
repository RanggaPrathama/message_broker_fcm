package utils

import (
	"time"

	"github.com/RanggaPrathama/message_broker_fcm/domain/models"
	"github.com/RanggaPrathama/message_broker_fcm/lib"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(idUser uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
		UserId: idUser,
		Email:  username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	tokenString, err := token.SignedString([]byte(lib.LoadEnv("SECRET_KEY")))

	if err != nil {
		return "token undifined", err
	}

	return tokenString, nil
}