package crypt

import "testing"

func TestEncrypt(t *testing.T) {
	plainText := "This is a plain text data"
	secretKey := "NotSoSecret"

	encryptedData, _ := Encrypt([]byte(plainText), secretKey)

	if encryptedText := string(encryptedData); encryptedText == plainText {
		t.Errorf("Encrypted data shouldn't be equals to the plain text data: (expected %s != %s)", plainText, encryptedText)
	}
}

func TestDecrypt(t *testing.T) {
	plainText := "This is a plain text data"
	secretKey := "NotSoSecret"

	encryptedData, _ := Encrypt([]byte(plainText), secretKey)
	decryptedData, _ := Decrypt(encryptedData, secretKey)

	if decryptedText := string(decryptedData); decryptedText != plainText {
		t.Errorf("Decrypted data should be equals to the plain text data: (expected %s = %s)", plainText, decryptedText)
	}
}

func TestDecryptFailureWithDifferentSecret(t *testing.T) {
	plainText := "This is a plain text data"
	secretKey := "NotSoSecret"

	encryptedData, _ := Encrypt([]byte(plainText), secretKey)
	_, err := Decrypt(encryptedData, "DifferentSecretKey")

	if err == nil {
		t.Errorf("Expected decrypt to return an error because the secretKey from encrypt and decrypt are different.")
	}
}

func BenchmarkEncrypt(b *testing.B) {
	plainText := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec at felis lorem. 
	Quisque elementum sem in erat tristique, id ultricies neque sollicitudin. Vivamus malesuada finibus eros, 
	vitae tristique nisl facilisis vitae. In tempor neque a posuere blandit. Etiam pulvinar orci ac nisi gravida maximus. 
	Curabitur id tincidunt augue. Morbi sollicitudin dolor risus. Phasellus ultricies diam non nunc molestie rhoncus. 
	Cras sollicitudin nibh at sapien tristique, at laoreet diam convallis.`

	// Creates a relative bigger text to be encrypted, using the concat with itself ten times
	for i := 0; i < 10; i++ {
		plainText += plainText
	}

	secretKey := "}#?~f6^}K/$E-T#Bg*g!Kd-V#fZ'}nJe"

	for n := 0; n < b.N; n++ {
		Encrypt([]byte(plainText), secretKey)
	}
}

func BenchmarkEncryptDecrypt(b *testing.B) {
	plainText := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec at felis lorem. 
	Quisque elementum sem in erat tristique, id ultricies neque sollicitudin. Vivamus malesuada finibus eros, 
	vitae tristique nisl facilisis vitae. In tempor neque a posuere blandit. Etiam pulvinar orci ac nisi gravida maximus. 
	Curabitur id tincidunt augue. Morbi sollicitudin dolor risus. Phasellus ultricies diam non nunc molestie rhoncus. 
	Cras sollicitudin nibh at sapien tristique, at laoreet diam convallis.`

	// Creates a relative bigger text to be encrypted, using the concat with itself ten times
	for i := 0; i < 10; i++ {
		plainText += plainText
	}

	secretKey := "}#?~f6^}K/$E-T#Bg*g!Kd-V#fZ'}nJe"

	for n := 0; n < b.N; n++ {
		encryptedData, _ := Encrypt([]byte(plainText), secretKey)
		Decrypt(encryptedData, secretKey)
	}
}
