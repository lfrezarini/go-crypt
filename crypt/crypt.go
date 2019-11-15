package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func Encrypt(data []byte, secret string) []byte {
	passphrase := []byte(CreatePassphrase(secret))
	block, err := aes.NewCipher(passphrase)

	if err != nil {
		log.Panicf("Error while trying to create a cipher block: %v", err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		log.Panicf("Error while creating the Galois Counter: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Panicf(err.Error())
	}

	return gcm.Seal(nonce, nonce, data, nil)
}

func Decrypt(data []byte, secret string) ([]byte, error) {
	passphrase := []byte(CreatePassphrase(secret))
	block, err := aes.NewCipher(passphrase)

	if err != nil {
		return nil, fmt.Errorf("Error while trying to create a cipher block: %v", err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, fmt.Errorf("Error while creating the Galois Counter: %v", err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		return nil, fmt.Errorf("Error while trying to decrypt the data: %v", err)
	}

	return plaintext, nil
}
