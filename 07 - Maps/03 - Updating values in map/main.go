package main

import "fmt"

func main() {

	menu := map[string]string{
		"soup":    "Chicken",
		"Salad":   "Ceaser",
		"Starter": "Veg",
		"Dessert": "Cold",
	}

	fmt.Printf("Old value: %v \n",menu["soup"])
	menu["soup"] = "Mutton"
	fmt.Printf("New value: %v \n", menu["soup"])
}