package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to string parsing")

	teststring := "This is a test string"

	// Split at each character
	teststringparsed := strings.Split(teststring, "")
	fmt.Println(teststringparsed[0])

	// Split at white spaces
	teststringparsed = strings.Split(teststring," ")
	fmt.Println(teststring)

	strings.
}