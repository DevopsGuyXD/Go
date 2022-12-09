package main

import "fmt"

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good bye %v \n", n)
}

func cycleNames(n []string, f func(string)){
	for _,v := range n {
		f(v)
	}
}

func main() {
	sayGreeting("Mario")
	sayBye("Luigi")

	cycleNames([]string{"Luigi", "Mario", "Peach", "Bowser"}, sayGreeting)
}