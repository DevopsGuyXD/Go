package main

import "github.com/devopsguyXD/test/02-Abstraction/handler"

func main() {
	dog := handler.NewCreature("Brown", 4, "Bite")
	handler.Attack(dog)
}