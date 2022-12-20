package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	n := 0 
	for k, v := range menu {

		fmt.Printf(" %v. You key is %q and your value is %v \n", n, k, v)		
	}

}