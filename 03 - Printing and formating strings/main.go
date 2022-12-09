package main

import "fmt"

func main() {

	// Same line
	fmt.Print("Hello world")
    
	// New Line
	fmt.Println("Hello world")


	// Parsing variable to strings
	name := "Mario"
	age := 30

	fmt.Println("My name is",name,"and my age is",age)

	//Formatted string
	fmt.Printf("My name is %v and my age is %v \n", name, age)

	//Format with quotes around string
	fmt.Printf("My name is %v and my age is %q \n", name, age)
}