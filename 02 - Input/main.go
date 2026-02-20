package main

import (
	"bufio"
	"fmt"
	"os"
)

// This method is useful if your input when prompted is a single word or number
func singleValue() {
	var response string
	fmt.Println("What is the rating of the pizza:")
	fmt.Scanln(&response)
  
	fmt.Printf("You responded with %v\n", response)
}

// This method is useful if your input when prompted has multiple words 
func multiValue() {
	reader := bufio.NewReader(os.Stdin) 
	fmt.Println("What is the rating of the pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}

func main() {
	singleValue()
	multiValue()
}
