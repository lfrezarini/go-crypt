# Go Crypt

Go crypt is a simple CLI written in Go that allows you to encrypt and decrypt files. It was made by educational proposes and shouldn't be used in production environments.

## Instalation

Of course, in order to test this application, you should have Go installed in your machine. [here](https://golang.org/dl/) you can download and follow the instructions to install Go on your machine.

After that, you can download this project using `go get` command or clonning this repository using git.

### Using go get

Run the following command:

```sh
go get -u github.com/LucasFrezarini/go-cyrpt
```

### Using GIt

First, you have to clone this repository:

```sh
git clone https://github.com/LucasFrezarini/go-cyrpt.git
```

after that, you'll have to install the project dependencies using `go mod`:

```sh
go mod download
```

then, you can run the CLI using:

```sh
go run main.go [command] [flags]
```