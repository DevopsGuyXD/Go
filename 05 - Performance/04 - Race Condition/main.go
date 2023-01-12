package main

import "fmt"

func main() {

	fmt.Println("Race condition")

	var score = []int{0}

	func(){
		fmt.Println("One R")
		score = append(score, 1)
	}()
}