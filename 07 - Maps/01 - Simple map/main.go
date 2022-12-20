package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	// Print the entire Map
	fmt.Println(menu)

	// Print the value of a specific key
	fmt.Println(menu["pie"])

}