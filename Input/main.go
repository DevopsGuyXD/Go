package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is the rating of the pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
