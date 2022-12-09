package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname,p.lname, `says, "Good moring, james."`)
}

func main() {

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	p1.speak()
}