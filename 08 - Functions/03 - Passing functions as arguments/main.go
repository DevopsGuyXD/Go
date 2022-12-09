package main

import "fmt"

func sayGreeting(n string) {
	fmt.Printf("Good morning %q \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good bye %q \n", n)
}

func cycleNames(f1 func(string), f2 func(string)){
	f1("Mario")
	f2("Peach")
}

func main() {

	cycleNames(sayGreeting, sayBye)
}