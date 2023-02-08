package main

import (
	"fmt"
	"net/http"
	"os"

	router "github.com/Nowfloats/Bizapp/Routers"
	util "github.com/Nowfloats/Bizapp/Utils"
)

func main() {
	
	util.InitEnvFile()

	server := router.RouterCollection()
	fmt.Println("Server listening on port 8000")

	err := http.ListenAndServe(os.Getenv("PORT"), server); util.CheckForNil(err)
}