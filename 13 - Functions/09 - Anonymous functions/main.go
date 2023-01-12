package main

import (
	"fmt"
)

func main() {
	
	// Simple function
	func(){
		fmt.Println("Welcome to anonymous functions")
	}()


	// Simple function with parameters
	func(i float64){
	
		multi := 34.36
	
		multi = multi * i
	
		fmt.Println(multi)
	}(67.7)


	// Simple function assigned to a variable
	var greeting = func(){
		fmt.Println("Hello from Mars")
	}
	greeting()

	// Simple fucntion with return 
	var sum = func(n1 int, n2 int) int{

		sum := n1 + n2
		return sum
	}

	result := sum(5, 4)
	fmt.Print(result)
}