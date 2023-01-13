package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("Welcome to URL parsing")

	testurl := "http://www.sitemaps.org/schemas/sitemap/0.9"

	testurlparsed, _ := url.Parse(testurl)

	fmt.Println(testurlparsed.Scheme)
	fmt.Println(testurlparsed.Host)
	fmt.Println(testurlparsed.Path)
}