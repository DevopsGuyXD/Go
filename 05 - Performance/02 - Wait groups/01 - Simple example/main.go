package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup


func main(){

	wg.Add(2)

	mark1()
	mark2()

	wg.Wait()

}

func mark1() {

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Mark1")
	}

	wg.Done()
}

func mark2() {

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Mark2")
	}

	wg.Done()
}