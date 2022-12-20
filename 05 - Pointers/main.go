package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")

	myNumber := 23

	var ptr = &myNumber

	fmt.Println(ptr)
	fmt.Println(*ptr)
}