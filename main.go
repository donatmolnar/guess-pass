package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt" //go get golang.org/x/crypto/bcrypt
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func passcheck(w http.ResponseWriter, r *http.Request) {
	password := "secret"
	hash, _ := HashPassword(password)
	match := CheckPasswordHash(r.URL.Path[1:], hash)
	fmt.Fprintf(w, "Provided password: %s\n", r.URL.Path[1:])
	fmt.Fprintf(w, "Password has matched: %v\n", match)
	if match {
		fmt.Fprintf(w, "You have guessed right!")
	} else {
		fmt.Fprintf(w, "Wrong! Not even close...")
	}
}

func main() {
	http.HandleFunc("/", passcheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
