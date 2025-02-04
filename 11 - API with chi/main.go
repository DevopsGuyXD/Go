package main

import (
	"fmt"
	"net/http"

	router "github.com/DevopsGuyXD/testing/Routes"
)

type Test struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {

	server := router.RouteCollection()

	err := http.ListenAndServe(":8000", server)
	if err != nil {
		fmt.Println(err)
	}
}
