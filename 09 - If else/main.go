package main

import "fmt"

func main() {

	fmt.Println("If else in golang:")

	age := 18
	var message string

	if age < 18 {
		message = "You are not of legal age"
	} else if age >= 18{
		message = "You are legal to drink"
	}

	fmt.Println(message)
}