package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	greeting := "Hello world"

	// Below are a few commmon examples
	fmt.Println(strings.Contains(greeting, "Hello"));

	fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi"));

	fmt.Println(strings.ToUpper(greeting));

	fmt.Println(strings.Index(greeting,"w"));

	fmt.Println(strings.Split(greeting," "));

	ages := []int{45,34,76,68,23,8,4,22,74,35,75,3}

	sort.Ints(ages)
	fmt.Println(ages);

	index := sort.SearchInts(ages, 34)
	fmt.Print(index)
}