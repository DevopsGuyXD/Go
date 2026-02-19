package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")

	age := 30

	adultYears := getAdultYears(&age)
	fmt.Println(adultYears)
	updateValueinMem(&age)
}

func getAdultYears(age *int) int {
	return *age - 10
}

func updateValueinMem(age *int){
	*age = *age - 10
}

// Use pointers only when values are very large
// Most often causes confusion if not used carefuly