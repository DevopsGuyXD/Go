package main

import "fmt"

type Animal interface{
	Attack()
}

type Dog struct {
	Color       string
	Limbs       int
	Attack_type string
}

type Cat struct {
	Color       string
	Limbs       int
	Attack_type string
}

func (d Dog) Attack() string {
	attack_points := 40
	result := fmt.Sprintf("Attack type: %v\nAttack points: %v", d.Attack_type, attack_points)
	return result
}

func (c Cat) Attack() string {
	attack_points := 5
	result := fmt.Sprintf("Attack type: %v\nAttack points: %v", c.Attack_type, attack_points)
	return result
}

func main() {

	dog := Dog{Color: "Black", Limbs: 4, Attack_type: "Bite"}
	cat := Cat{Color: "Brown", Limbs: 4, Attack_type: "Scratch"}
	
	fmt.Println(dog.Attack())
	fmt.Println(cat.Attack())

}