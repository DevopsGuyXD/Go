package main

import "fmt"

func multiFunc(values ...int) (int, string) {

	var total int

	for _, value := range values {
		total += value
	}

	return total, "Hi pro result function"
}

func main() {
	multi, message := multiFunc(4, 5, 6, 7)

	fmt.Println(multi)
	fmt.Println(message)

}