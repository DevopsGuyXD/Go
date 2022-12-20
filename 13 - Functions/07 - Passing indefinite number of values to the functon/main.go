package main

import "fmt"

func indefiniteFun(values ...int) int {

	var total int

	for _, value := range values {
		total += value
	}

	return total
}

func main() {
	indef := indefiniteFun(4, 6, 7, 5, 8)
	fmt.Printf("%v",indef)
}