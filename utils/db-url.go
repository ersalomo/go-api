package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetDbUrl() (string, string) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error load .env file", err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	//dbRootPassword := os.Getenv("DB_ROOT_PASSWORD")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, "localhost", "3306", database)
	return dbDriver, dbUri
}
