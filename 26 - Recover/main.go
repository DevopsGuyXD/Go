package main

import (
	"fmt"
)

func main() {
	panicExample()
	fmt.Println("Post recovery")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func panicExample() {
	defer recoverFromPanic()
	panic("Oops! Something went wrong.")
}
