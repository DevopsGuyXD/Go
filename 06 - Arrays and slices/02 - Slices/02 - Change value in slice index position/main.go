package main

import "fmt"

func main() {

	scores := []int{100, 42, 56}
	scores[2] = 57

	fmt.Println(scores)

}