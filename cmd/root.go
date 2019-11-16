package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// FilePath determines the file to be encrypted/decrypted
	FilePath string
	// Secret determines the secret word to be used on encryption/decryption process
	Secret string
	// Output determines the file to write the result of encryption/decryption process
	Output string
)

var rootCmd = &cobra.Command{
	Use:   "go-crypt",
	Short: "Go-crypt is a utility-tool to easily encrypt and decrypt files using AES.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&FilePath, "file", "f", "", "Path to file that should be encrypted/decrypted")
	rootCmd.PersistentFlags().StringVarP(&Secret, "secret", "s", "", "Secret to be used to encrypt/decrypt the file")
	rootCmd.PersistentFlags().StringVarP(&Output, "output", "o", "", "Output to write the results of the encrypting/decrypting (defaults to OS stdout)")

	rootCmd.MarkPersistentFlagRequired("file")
	rootCmd.MarkPersistentFlagRequired("secret")
}

// Execute starts the Cobra CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
