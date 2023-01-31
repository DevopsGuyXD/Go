package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var test_map map[string]string
	test_string := `{"EMAIL_SENDER":"nfcloudsecurity@nowfloats.com","EMAIL_SENDER_PASSWORD":"P67%4urG123"}`

	json.Unmarshal([]byte(test_string), &test_map)

	fmt.Printf("%T \n" ,test_string)
	fmt.Printf("%T \n" ,test_map)
}