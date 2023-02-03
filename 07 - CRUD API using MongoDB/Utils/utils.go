package utils

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitEnvFile(){
	err := godotenv.Load(".env")
	CheckForNil(err)
}

func FormatedTime() string{
	currentTime := time.Now().Format("15:04:05 02-01-2006")

	return currentTime
}