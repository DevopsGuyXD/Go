package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Course struct {
	Name     string `json:"name"`
	Price    int	`json:"price"`
	Platform string `json:"platform"`
	Password string	`json:"-"`
}

func main() {

	EncodeJson()
}

func EncodeJson() {
	course := []Course{
		{"Java", 10, "Linux", "Pa55word"},
		{"Dotnet", 7, "Windows", "Password"},
	}

	finalJson, err := json.MarshalIndent(course,"","\t"); if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(finalJson))
}