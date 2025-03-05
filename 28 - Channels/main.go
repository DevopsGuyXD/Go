package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func fetchAPIChannel(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	res, _ := http.Get(url)
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	ch <- fmt.Sprintf("Concurrent Length: %d", len(body))
}

func fetchAPI(url string) string {
	res, _ := http.Get(url)
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	return fmt.Sprintf("Sequential Length: %d", len(body))
}

func main() {
	fmt.Println("Goroutines and Channels test")

	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/4",
		"https://jsonplaceholder.typicode.com/todos/5",
	}

	//----- Sequential
	start := time.Now()
	for _, url := range urls {
		fmt.Println(fetchAPI(url))
	}
	fmt.Printf("Sequential Total:%v \n\n", time.Since(start))

	//----- Concurrent
	start = time.Now()
	var wg sync.WaitGroup
	ch := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchAPIChannel(url, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		fmt.Println(res)
	}

	fmt.Printf("Concurrent Total:%v \n\n", time.Since(start))

}
