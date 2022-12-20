package main

import "fmt"

func testFunc() {

	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	testFunc()
}