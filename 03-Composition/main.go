package main

import "fmt"

type Dog struct {
	Animal
}

type Animal struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "barks")
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}
 
func main() {
	d := Dog{Animal{Name: "Buddy"}}
	d.Speak()
	d.Animal.Speak()
}
