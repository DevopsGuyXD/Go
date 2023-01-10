package main

import (
	"fmt"
	"net/http"
	"os"

	config "github.com/DevopsguyXD/test/Config"
	routes "github.com/DevopsguyXD/test/Routes"
	utils "github.com/DevopsguyXD/test/Utils"
)

func main() {

	utils.InitEnvFile()
	config.MongoConnection()

	server := routes.RouteCollection()

	fmt.Println("Server listening on port 8000")

	err := http.ListenAndServe(os.Getenv("PORT"),server)
	utils.CheckForNil(err)
}