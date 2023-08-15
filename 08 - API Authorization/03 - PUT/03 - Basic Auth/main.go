package main

import (
	"net/http"

	config "github.com/DevopsGuyXD/testserver/Configs"
	route "github.com/DevopsGuyXD/testserver/Routes"
)

func main() {

	config.MongoConnection()

	server := route.RouteCollection()

	http.ListenAndServe(":8001", server)
}