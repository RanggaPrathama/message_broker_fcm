package utils

import (
	"time"

	"github.com/RanggaPrathama/message_broker_fcm/lib"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(idUser uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user": idUser,
		"username": username,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(lib.LoadEnv("SECRET_KEY")))

	if err != nil {
		return "token undifined", err
	}

	return tokenString, nil
}