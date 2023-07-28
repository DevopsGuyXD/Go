package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main(){

	formData := url.Values{}
	formData.Add("DOMAIN_NAME", "www.google.com")
    formData.Add("API_CALL_TYPE", "GET")
    formData.Add("CHECK_INTERVALS", "5m")

	req, err := http.NewRequest("POST", "http://localhost:8000/add_domain", strings.NewReader(formData.Encode())); util.CheckForNil(err)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	resp, err := client.Do(req); util.CheckForNil(err)

	defer resp.Body.Close()

	fmt.Println("Sent successfully")
}