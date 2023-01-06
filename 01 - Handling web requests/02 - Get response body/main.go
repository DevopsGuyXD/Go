package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("Geting the response body")

	url := "https://www.google.com"

	response, err := http.Get(url); if err != nil{
		panic(response)
	}else{
		response_body, err := ioutil.ReadAll(response.Body); if err != nil{
			panic(err)
		}else{
			fmt.Println(string(response_body))
		}
	}
}