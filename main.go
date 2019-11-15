package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LucasFrezarini/go-crypt/crypt"
)

func main() {
	args := os.Args[1:]

	fmt.Printf("Encrypting %v...\n", args)
	result := crypt.Encrypt([]byte(args[0]), "calopsita")

	decrypted, err := crypt.Decrypt(result, "calopsita")

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(result), string(decrypted))
}
