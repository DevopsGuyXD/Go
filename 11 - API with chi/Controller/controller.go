package controller

import (
	"encoding/json"
	"net/http"
)

func Test1Controller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Working fine - Test 1")
}

func Test2Controller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Working fine - Test 2")
}
