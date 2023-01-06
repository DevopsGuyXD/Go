package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env"); if err != nil{
		panic("Cannot load .env file")
	}

	fmt.Println(err)

	fmt.Println(os.Getenv("MESSAGE"))
}