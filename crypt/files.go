package crypt

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
)

// EncryptFileContent encrypts a file, line by line, with the AES algorithm, creathing a passphrase with the provided secret
func EncryptFileContent(path, secret string, output io.Writer) error {
	f, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("Error while opening the file: %v", err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadBytes('\n')

		if err != nil {
			if err == io.EOF {
				str := base64.StdEncoding.EncodeToString(Encrypt(line, secret))
				_, err = output.Write([]byte(str))

				if err != nil {
					return fmt.Errorf("Error while writing the output: %v", err)
				}

				break
			}

			return fmt.Errorf("Error while reading the file: %v", err)
		}

		str := base64.StdEncoding.EncodeToString(Encrypt(line, secret))

		_, err = output.Write([]byte(str + "\n"))

		if err != nil {
			return fmt.Errorf("Error while writing the output: %v", err)
		}
	}

	return nil
}

// DecryptFileContent decrypts a file that was encrypted with the Base64 + AES algorithm with the same secret
func DecryptFileContent(path, secret string, output io.Writer) error {
	f, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("Error while opening the file: %v", err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		b64Line, err := r.ReadBytes('\n')

		if err != nil {
			if err == io.EOF {
				line, _ := base64.StdEncoding.DecodeString(string(b64Line))
				data, err := Decrypt(line, secret)
				log.Print(data, err)
				output.Write(data)
				break
			}

			return fmt.Errorf("Error while reading the file: %v", err)
		}

		line, _ := base64.StdEncoding.DecodeString(string(b64Line))
		data, err := Decrypt(line, secret)

		if err != nil {
			return fmt.Errorf("Error while decrypting line: %v", err)
		}

		_, err = output.Write(data)

		if err != nil {
			return fmt.Errorf("Error while writing to the file: %v", err)
		}
	}

	return nil
}
