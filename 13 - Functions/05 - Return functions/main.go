package main

import (
	"fmt"
	"math"
)

func circleArea(r float64) float64 {
	return math.Pi*r*r
}

func main() {
	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Printf("Your value is %0.3f \n",a1)
	fmt.Printf("Your value is %0.3f \n",a2)
}