package lib

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return os.Getenv(name)
}

	