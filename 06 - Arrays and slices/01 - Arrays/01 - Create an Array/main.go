package main

import "fmt"

func main() {

	// Long hand
	var firstAges [3]int = [3]int{20, 25, 35}

	// Short hand
	var secondAges = [3]int{20, 25, 35}

	// Shorter hand
	thirdAges := [3]string{"Yoshi", "Mario", "Luigi"}

	fmt.Println(firstAges);
	fmt.Println(secondAges);
}