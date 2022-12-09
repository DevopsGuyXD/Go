package main

import "fmt"

func sayGreeting(n string) {
	fmt.Printf("Good morning to you %q \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good bye to you %q \n", n)
}

func main() {

	sayGreeting("Mario")
	sayGreeting("Peach")
	sayBye("Bowser")
}