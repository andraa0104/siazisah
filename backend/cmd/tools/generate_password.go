package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/tools/generate_password.go <password>")
		return
	}

	password := os.Args[1]
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Println(string(hash))
}
