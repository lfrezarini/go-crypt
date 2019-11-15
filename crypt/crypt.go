package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Encrypt encrypts the data using the AES algorithm, creating a passphrase with the provided secret
func Encrypt(data []byte, secret string) ([]byte, error) {
	passphrase := []byte(CreatePassphrase(secret))
	block, err := aes.NewCipher(passphrase)

	if err != nil {
		return nil, fmt.Errorf("Error while trying to create a cipher block: %v", err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, fmt.Errorf("Error while creating the Galois Counter: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// Decrypt Decrypt the data using the AES algorithm, using the secret as key to the passphrase
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
