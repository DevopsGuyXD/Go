package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Platform string `json:"platform"`
	Password string `json:"-"`
}

func main() {
	DecodeJson()
}

func DecodeJson() {
	jsonData := []byte(`
		{
		"coursename": "ReactJS",
		"Price":      299,
		"website": "google.com"
		}
	`)

	var courses Course

	checkValid := json.Valid(jsonData)

	if checkValid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &courses)
		fmt.Printf("%#v\n", courses)
	}else{
		fmt.Println("JSON was not valid")
	}
}