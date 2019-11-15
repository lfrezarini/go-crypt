package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

// CreatePassphrase creates a new 32-byte hash that can be used as a passphrase on encrypting algoritms.
func CreatePassphrase(secret string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum(nil))
}
