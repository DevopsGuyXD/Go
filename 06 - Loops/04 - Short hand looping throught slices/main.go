package main

import "fmt"

func main() {
	names := []string{"Mario", "Luigi", "Bowser", "Peach"}

	// With index
	for index, value := range names {
		fmt.Printf("%v. Your name is %v \n", index, value)
	}

	//Without index
	for _,value := range names{
		fmt.Printf("Your name is %v \n", value)
	}
}