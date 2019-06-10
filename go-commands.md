# Go useful commands explained

Here are a list of *daily Go commands* used by developers, which will
help you stay productive if you know about them.

## Frequently used commands

### `go get`

Downloads *third party packages*, by *git* cloning the repository
and installing binaries if any

Example:
```bash
# downloads the "errors" package
go get github.com/pkg/errors

# downloads and updates the "errors" package
go get -u github.com/pkg/errors

# downloads all packages that the current project imports/uses
go get ./...
```

### `go run`

Compiles and builds the project, after which it generates a binary in
a temporary location then executes that binary

Example:
```bash
# Compiles "main.go" and executes the binary
go run main.go

# Compiles "main.go" and "package.go" and executes the binary
go run main.go package.go

# Compiles all ".go" files in CWD and executes the binary
go run *.go
```

### `go build`

Compiles and creates executable binary from source code

Example:
```bash
go build

go build -o executable-name
```

### `go test`

Runs all Go tests which are located inside *_test.go* files from
current working directory

Example:
```bash
go test ./...
```

## Other useful commands

### `go install`

Compiles Go code from binary.go and creates a binary called
"bin" then places it inside `$GOPATH/bin` directory

In order for this to work the file need to in **package main**
and have a **main** function

Example:
```bash
go install bin.go
```

### `go fmt`

Formats all Go code from current working directory

Example:
```bash
go fmt ./...
```

### `go vet`

Checks for potential Go issues in current working directory

Example:
```bash
go vet ./...
```

### `go env`

Displays all environment variables that Go uses and their values

### `go help`

Displays a list of all available *Go commands*. Have fun and play
around with them

For more info about Go commands check out
[Go Commands](https://golang.org/cmd/go/)

Back to
[Go Basics](https://github.com/gophertuts/go-basics)
