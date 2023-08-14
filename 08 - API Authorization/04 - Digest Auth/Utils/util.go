package util

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	username = "bharathdundi"
	password = "Pa55word@123"
	realm    = "MyRealm"
	nonce    = "1234567890"
    opaque   = "xyz"
)

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}





// Digest
func sendDigestChallenge(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Digest realm="%s", qop="auth", nonce="%s", opaque="%s"`, realm, nonce, opaque))
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}





func validateDigest(authHeader string, method string) bool {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Digest" {
		return false
	}

	authInfo := make(map[string]string)
	for _, part := range strings.Split(parts[1], ",") {
		parts := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(parts) == 2 {
			key := strings.Trim(parts[0], `"`)
			value := strings.Trim(parts[1], `"`)
			authInfo[key] = value
		}
	}

	expectedResponse := generateExpectedResponse(method, authInfo["uri"], authInfo["nc"], authInfo["cnonce"])

	// Compare the calculated response with the received response
	return authInfo["response"] == expectedResponse
}

func generateExpectedResponse(method, uri, nc, cnonce string) string {
	ha1 := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", username, realm, password))))
	ha2 := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s", method, uri))))
	response := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, nonce, nc, cnonce, "auth", ha2))))
	return response
}







func DigestAuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			sendDigestChallenge(w)
			return
		}

		if !validateDigest(authHeader, r.Method) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}