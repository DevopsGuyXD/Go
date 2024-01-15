package main

import "fmt"

type Bill interface{
	PrintPrice()
}

type Book struct {
	Name  string
	Price int
}

type Drink struct {
	Name  string
	Price int
}

type Puzzle struct {
	Name  string
	Price int
}

func (b Book) PrintPrice(){
	fmt.Printf("Name: %v | Price: %v\n", b.Name, b.Price)
}

func (d Drink) PrintPrice(){
	fmt.Printf("Name: %v | Price: %v\n", d.Name, d.Price)
}

func (p Puzzle) PrintPrice(){
	fmt.Printf("Name: %v | Price: %v\n", p.Name, p.Price)
}

func main() {
	fmt.Println("Welcome to the interface test")

	book := Book{"Java", 20}
	drink := Drink{"Coffee", 10}
	puzzle := Puzzle{"Rubics cube", 5}

	info := []Bill{ book, drink, puzzle }

	info[0].PrintPrice()
	info[1].PrintPrice()
	info[2].PrintPrice()
}