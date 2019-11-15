package cmd

import (
	"io"
	"log"
	"os"

	"github.com/LucasFrezarini/go-crypt/crypt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file encrypted using the AES algorithm + base64",
	Run: func(cmd *cobra.Command, args []string) {
		var output io.Writer
		var err error

		if Output == "" {
			output = os.Stdout
		} else {
			output, err = os.Create(Output)

			if err != nil {
				log.Fatalf("Error while trying to create the output file: %v", err)
			}
		}

		if err = crypt.DecryptFileContent(FilePath, Secret, output); err != nil {
			log.Fatalf("Error while trying to decrypt the file content: %v", err)
		}

		if Output != "" {
			log.Printf("Successful decrypted into %s\n", Output)
		}
	},
}
