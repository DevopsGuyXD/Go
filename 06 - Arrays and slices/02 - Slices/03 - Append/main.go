package main

import "fmt"

func main() {
	scores := []int{10, 45, 78, 34}
	scores = append(scores, 67)

	fmt.Println(scores)
}