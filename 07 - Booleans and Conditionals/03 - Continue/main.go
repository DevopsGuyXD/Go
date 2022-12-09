package main

import "fmt"

func main() {

	names := []string{"Mario", "Luigi", "Bowser", "Peach"}

	for index, value := range names {

		if index == 1 {
			fmt.Printf("%v. Continuing at position %q \n", index, value)
			// Go back to top of the loop
			continue
		}

		fmt.Printf("%v. Your name is %q \n", index, value)
	}
}