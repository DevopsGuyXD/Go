package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var test_map map[string]string
	test_string := `{"EMAIL_SENDER": <EMAIL ID>,"EMAIL_SENDER_PASSWORD": <PASSWORD>}`

	json.Unmarshal([]byte(test_string), &test_map)

	fmt.Printf("%T \n" ,test_string)
	fmt.Printf("%T \n" ,test_map)
}