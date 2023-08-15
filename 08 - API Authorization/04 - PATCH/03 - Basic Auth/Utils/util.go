package util

import (
	"log"
	"net/http"
)

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        username, password, ok := r.BasicAuth()

        if !ok || !isValidUser(username, password) {
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        handler(w, r)
    }
}

func isValidUser(username, password string) bool {

	var validUsers = map[string]string{
		"bharathdundi": "Pa55word@123",
	}

    storedPass, exists := validUsers[username]
    if !exists || storedPass != password {
        return false
    }
    return true
}