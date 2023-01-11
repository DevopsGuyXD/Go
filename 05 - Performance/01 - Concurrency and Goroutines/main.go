package main

import (
	"fmt"
	"time"
)

func main() {
	go Greeter("Hello")
	Greeter("World")
}

func Greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}