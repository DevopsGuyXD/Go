package main

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Age       int
}

func (u User) GetFirstname() string {
	return u.Firstname
}

func (u User) GetLastname() string {
	return u.Lastname
}

func main() {

	user := User{Firstname: "Goldie", Lastname: "Dundi", Age: 9}

	fmt.Println(user.GetFirstname())
	fmt.Println(user.GetLastname())
}