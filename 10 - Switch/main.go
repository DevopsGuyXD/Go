package main

import "fmt"

func main() {
	fmt.Println("Switch and case in golang")

	num := 8

	switch num{
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6")
	default:
		fmt.Println("Rolled off the board")
	}
}