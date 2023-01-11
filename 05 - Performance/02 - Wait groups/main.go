package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.getboost360.com",
		"https://www.Netflix.com",
		"https://www.Ibanez.com",
	}

	go getURLS(websiteList)
	go testfunc()
	
	wg.Add(2)
	wg.Wait()
}

func getURLS(url []string) {
	defer wg.Done()
	for _, value := range url{
		res, err := http.Get(value); if err !=nil{
			log.Fatal(err)
		}

		fmt.Println(res.StatusCode ,value)
	}
}

func testfunc(){

	defer wg.Done()
	for i := 0; i <= 2000; i++{
		fmt.Println(i)
	}
}