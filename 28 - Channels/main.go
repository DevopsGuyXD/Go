package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Welcome to channels")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)

		go func(ch <-chan int, wg *sync.WaitGroup){
			fmt.Println(<- myCh)
			wg.Done()
		}(myCh, wg)

		go func(ch chan<- int, wg *sync.WaitGroup){
			myCh <- 5
			close(myCh)
			wg.Done()
		}(myCh, wg)

	wg.Wait()

}