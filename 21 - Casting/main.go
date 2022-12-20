package main

import (
	"fmt"
	"strconv"
)

func main() {

	num1 := "1"

	num1_conv, err := strconv.ParseFloat(num1, 64); if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Converted num1 to float")
		num3 := num1_conv + 1
		fmt.Println(num3)
	}
}