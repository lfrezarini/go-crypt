# Go Crypt

Go crypt is a simple CLI written in Go that allows you to encrypt and decrypt files. It was made by educational proposes and shouldn't be used in production environments.

## Instalation

Of course, in order to test this application, you should have Go installed in your machine. [here](https://golang.org/dl/) you can download and follow the instructions to install Go on your machine.

After that, you can download this project using `go get` command or clonning this repository using git.

### Using go get

Running the following command, you'll be able to pull the project and run it as a binary in your system:

```sh
go get -u github.com/LucasFrezarini/go-crypt
go install github.com/LucasFrezarini/go-crypt
```

Then, you can execute it using either:

```sh
~/go/bin/go-crypt [command] [flags]
```

or, if you configured the $GOPATH/bin to be in your $PATH, you can just run as a normal command:

```sh
go-crypt [command] [flags]
```

### Using Git

First, you have to clone this repository:

```sh
git clone https://github.com/LucasFrezarini/go-crypt.git
```

after that, you'll have to install the project dependencies using `go mod`:

```sh
go mod download
```

then, you can run the CLI using:

```sh
go run main.go [command] [flags]
```

## Utilization guide

Example of usage:

```sh
go-crypt encrypt -f examples/plaintext.txt -o examples/encrypted -s SuperSecretKey
```

```sh
go-crypt decrypt -f examples/encrypted -o examples/decrypted.txt -s SuperSecretKey
```

There are two basic commands that you can use in this CLI: **encrypt** and **decrypt**. The encryption algoritm used is the AES, and the encryption process is made line by line, and every line is transformed to Base64, avoiding breaking the algoritm when the encryption generates line breaks. The result of the encrypt/decrypt process can be saved into a file or printed on the stdout.

## Flags

Either encrypt and decrypt commands have two mandatory flags: **--secret** and **--file**, and accepts an optional flag called **--output**:

## "--secret | -s" \<secret-key\>

That's the secret key that will be used to encrypt/decrypt each line of the file. The algorithm uses this key as a base to generate a 32-byte passphrase used in the encryption process. Of course, the same secret-key that you used to encrypt a file should be the one used to decrypt it. Trying to use a invalid key will result in an error.

## "--file | -f" \<path-to-file\>

This one is kinda obvious: this is the path to the file that you want to encrypt. If you pass an invalid/inexistent path, an error will be thrown.

## "--output | -o" \<output-file\>

This option allow you to save the encrypted data into a new file. If it is not provided, then the encrypted data will be printed on the system stdout.

## Commands

The two commands avaliable are self-explanatory: the **encrypt** and **decrypt** commands. In the examples showed above it was made clear how to use them too. Also, you can get more details about the commands and flags running:

```sh
go-crypt --help
```




