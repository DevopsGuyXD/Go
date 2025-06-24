package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./Controller/controller.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	_, err = file.WriteString("This is a test")
	if err != nil {
		fmt.Println("Error writing to file")
	}
}
