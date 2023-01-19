package utils

import (
	"encoding/json"
	"log"
)

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckIfValidJson(jsondata []byte) bool{
	checkValidJson := json.Valid(jsondata)

	return checkValidJson
}