package controller

import (
	"encoding/json"
	"net/http"
)

func Test1Controller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// w.Write([]byte("Working fine - Test 1")) // Use this if response is a non-JSON object
	json.NewEncoder(w).Encode("Working fine - Test 1")
}

func Test2Controller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Working fine - Test 2")
}

func TestQueryParameter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryParams := r.URL.Query()

	firstName := queryParams.Get("firstname")
	lastName := queryParams.Get("lastname")

	fullName := firstName + lastName

	json.NewEncoder(w).Encode(fullName)
}
