package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Handling simple web requests")

	url := "https://www.google.com"
	
	response , err := http.Get(url); if err != nil{
		panic(err)
	}else{
		fmt.Println(response)
		fmt.Println(response.Status)
		fmt.Println(response.ContentLength)
	}

	defer response.Body.Close()
}