package cmd

import (
	"io"
	"log"
	"os"

	"github.com/LucasFrezarini/go-crypt/crypt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a file using the AES algorithm",
	Run: func(cmd *cobra.Command, args []string) {
		var output io.Writer
		var err error

		if Output == "" {
			output = os.Stdout
		} else {
			output, err = os.Create(Output)
		}

		if err != nil {
			log.Fatalf("Error while trying to create the output file: %v", err)
		}

		if err = crypt.EncryptFileContent(FilePath, Secret, output); err != nil {
			log.Fatalf("Error while trying to encrypt the file content: %v", err)
		}

		if Output != "" {
			log.Printf("Successful encrypted into %s\n", Output)
		}
	},
}
