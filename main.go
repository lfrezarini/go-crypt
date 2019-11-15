package main

import (
	"log"
	"os"

	"github.com/LucasFrezarini/go-crypt/crypt"
)

func main() {
	secret := "calopsita"

	f, _ := os.Create("teste-encrypt.txt")
	err := crypt.EncryptFileContent("teste.txt", secret, f)

	if err != nil {
		log.Panic(err)
	}

	f.Close()

	f, _ = os.Create("teste-decrypt.txt")
	defer f.Close()

	err = crypt.DecryptFileContent("teste-encrypt.txt", secret, f)

	if err != nil {
		log.Panic(err)
	}
}
