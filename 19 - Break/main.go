package main

import "fmt"

func main() {

	names := []string{"Mario", "Luigi", "Bowser", "Peach"}

	for index, value := range names {

		if index == 1 {
			fmt.Printf("%v. Breaking at position %q \n", index, value)
			// Break out of the loop
			break
		}

		fmt.Printf("%v. Your name is %q \n", index, value)
	}
}