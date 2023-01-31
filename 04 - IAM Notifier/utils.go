package main

import (
	"log"

	"github.com/joho/godotenv"
)

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitEnvFile(){
	err := godotenv.Load(".env"); CheckForNil(err)
}