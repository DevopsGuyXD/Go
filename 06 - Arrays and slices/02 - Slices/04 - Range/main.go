package main

import "fmt"

func main() {
	scores := []int{20, 9, 65, 2, 78, 34}

	// From specifed index at start and finish
	rangeOne := scores[1:3]
	fmt.Println(rangeOne);

	// From specified to end
	rangeTwo := scores[2:]
	fmt.Println(rangeTwo);

	// From start till specified range
	rangethree := scores[:3]
	fmt.Println(rangethree)
}