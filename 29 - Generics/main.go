package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Manual way
func add[T int | float64](a T, b T) T {

	return a + b
}

// An easier way
func subtract[T constraints.Float | constraints.Integer](a T, b T) T {

	return a - b
}

func main() {

	integers_add := add(2, 2)
	floats_add := add(2.3, 5.2)

	integers_sub := subtract(5, 2)
	floats_sub := subtract(7.3, 5.2)

	fmt.Println(integers_add)
	fmt.Println(floats_add)

	fmt.Println(integers_sub)
	fmt.Println(floats_sub)
}
