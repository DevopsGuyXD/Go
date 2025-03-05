package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Function to make a simple API request
func fetchAPI(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	ch <- fmt.Sprintf("Concurrent Response Length: %d", len(body))
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/4",
		"https://jsonplaceholder.typicode.com/todos/5",
	}

	//---------------------------- Sequential Execution
	start := time.Now()
	for _, url := range urls {
		resp, _ := http.Get(url)
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("Sequential Response Length: %d\n", len(body))
	}
	fmt.Printf("Sequential Execution Time:%v\n\n", time.Since(start))

	////---------------------------- Concurrent Execution using Goroutines
	start = time.Now()
	var wg sync.WaitGroup
	ch := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchAPI(url, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		fmt.Println(res)
	}

	fmt.Println("Concurrent Execution Time:", time.Since(start))
}
