package crypt

import (
	"testing"
)

func TestCreatePassphrase(t *testing.T) {
	passphrase := CreatePassphrase("NotSoSecret")

	if size := len(passphrase); size != 32 {
		t.Errorf("Passphrase size must be 32, got %d", size)
	}
}
