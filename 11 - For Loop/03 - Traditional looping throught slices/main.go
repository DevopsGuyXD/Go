package main

import "fmt"

func main() {

	names := []string{"Mario", "Bowser", "Luigi", "Peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Your names is",names[i])
	}
}