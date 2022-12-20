package main

import "fmt"

func main() {

	menu := map[string]string{
		"soup":    "Chicken",
		"Salad":   "Ceaser",
		"Starter": "Veg",
		"Dessert": "Cold",
	}

	delete(menu, "soup")
	fmt.Println(menu)
}