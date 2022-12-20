package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	my_user := User{"Bharath", "bharathdundi@gmil.com", true, 16}

	fmt.Println(my_user)
	fmt.Printf("The details are: %+v\n", my_user)
	fmt.Printf("Name is: %v\n", my_user.Name)
}

type User struct{
	Name string
	Email string
	Status bool 
	Age int
}