package main

import "fmt"

func main() {

	// --------------------------------------------------------------------------------------------- Same line
	fmt.Print("Hello world")

	// --------------------------------------------------------------------------------------------- New Line
	fmt.Println("Hello world")

	name := "Mario"
	age := 30

	// --------------------------------------------------------------------------------------------- Parsing variable to strings
	fmt.Println("My name is", name, "and my age is", age)

	// --------------------------------------------------------------------------------------------- Formatted string
	fmt.Printf("My name is %v and my age is %v \n", name, age)

	// --------------------------------------------------------------------------------------------- Format with quotes around string
	fmt.Printf("My name is %v and my age is %q \n", name, age)

	// --------------------------------------------------------------------------------------------- Save formatted string to variable
	myString := fmt.Sprintf("My name is %v and my age is %q \n", name, age)
	fmt.Println(myString)

	// --------------------------------------------------------------------------------------------- Multi-Line
	multiLine := `This
					is
					an 
					exammple
					for
					multi
					line`
	print(multiLine)
}
