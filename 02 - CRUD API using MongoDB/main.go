package main

import (
	"fmt"
	"net/http"
	"os"

	controllers "github.com/DevopsguyXD/test/Controllers"
	routes "github.com/DevopsguyXD/test/Routes"
	utils "github.com/DevopsguyXD/test/Utils"
)

func main() {

	utils.InitEnvFile()
	controllers.MongoConnection()

	server := routes.RouteCollection()

	fmt.Println("Server listening on port 8000")

	err := http.ListenAndServe(os.Getenv("PORT"),server)
	utils.CheckForNil(err)
}