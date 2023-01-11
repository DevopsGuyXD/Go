package main

import (
	"fmt"
	"log"
	"net/http"

	routers "github.com/DevopsGuyXD/mongotest/Routers"
)

func main(){
	fmt.Println("Welcome to the MongoDB API")

	server := routers.Router()

	fmt.Println("Server starting...")
	err := http.ListenAndServe(":8000", server);if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Server listening on port 8000")
}