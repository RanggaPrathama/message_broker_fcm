package lib

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Database struct {
// 	db *gorm.DB
// }

var Database *gorm.DB

func ConnectionPostgree() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", 
	LoadEnv("DB_HOST"), 
	LoadEnv("DB_USER"),
	LoadEnv("DB_PASSWORD"),
	LoadEnv("DB_NAME"),
	LoadEnv("DB_PORT"),
	LoadEnv("DB_SSLMODE"),
	LoadEnv("DB_TIMEZONE"),
	) 
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	
	fmt.Println("Connection to database established")

	Database = db
	
	return Database
}