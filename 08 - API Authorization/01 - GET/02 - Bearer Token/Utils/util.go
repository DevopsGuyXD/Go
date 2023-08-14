package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte("ItJSGtVkoU8JB3QJgoZupPXenb5tuiEmpLh3EEIgYrM=")

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetBearerToken(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	claims := jwt.MapClaims{
		"sub": "user123",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iss": "your-issuer",
		"aud": "your-audience",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SigningKey); if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}
	
	json.NewEncoder(w).Encode(signedToken)
}


func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		authHeader := r.Header["Authorization"][0]

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return SigningKey, nil

		}); if err != nil{
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("not authorized: " + err.Error()))
        }

		if token.Valid{
			next(w, r)
		}

	})
}