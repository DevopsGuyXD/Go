package main

import (
	"fmt"
	"net/http"

	route "github.com/DevopsGuyXD/testserver/Routes"
)

func main() {

	fmt.Println("Test application")
	
	server := route.RouteCollection()
	http.ListenAndServe(":8001", server)
}