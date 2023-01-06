package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL construction")

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitest",
	}

	fmt.Println(partsOfUrl.String())
}