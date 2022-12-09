package main

import "fmt"

func main() {

	age := 10

	if age < 21 {
		fmt.Printf("You are %v. Not legal to drink \n", age)
	} else if age >=21 {
		fmt.Printf("You are %v. Legal to drink \n", age)
	} else {
		fmt.Println("Your age is not of this planet")
	}
}