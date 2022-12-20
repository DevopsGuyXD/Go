package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the lesson labels")

	goto my_label

	my_label:
		fmt.Print("This is my label")
}