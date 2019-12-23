# go-basics

## Go language basics

This section includes exercises and small projects which illustrate the basics of Go programming language. Things such as: types, data structures, functions, pointers, packages and others.

## Section overview

- [Go Learning Resources](https://github.com/gophertuts/go-basics/blob/master/go-learning-resources)
- [Go Program Anatomy](https://github.com/gophertuts/go-basics/blob/master/go-program-anatomy)
- [Number Systems](https://github.com/gophertuts/go-basics/blob/master/number-systems)
- [Basic Types](https://github.com/gophertuts/go-basics/blob/master/basic-types)
- [Constants](https://github.com/gophertuts/go-basics/blob/master/constants)
- [Packages](https://github.com/gophertuts/go-basics/blob/master/packages)
- [Vendor directory](https://github.com/gophertuts/go-basics/blob/master/vendor-directory)
- [Custom import paths](https://github.com/gophertuts/go-basics/blob/master/custom-import-paths)
- [Custom package manager](https://github.com/gophertuts/go-basics/blob/master/custom-package-manager)
- [Package Management](https://github.com/gophertuts/go-basics/blob/master/package-management)
- [Go modules](https://github.com/gophertuts/go-basics/blob/master/go-modules)
- [Go module proxy](https://github.com/gophertuts/go-basics/blob/master/go-module-proxy)
- [Internal directory](https://github.com/gophertuts/go-basics/blob/master/internal-directory)

### Installation

Before trying any of these examples make sure to have the `Go` binary installed on your platform.

OSX:

```bash
# Install Homebrew

/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

```bash
# Install Go

brew install go
```

Linux:

```bash
# Move into $HOME directory
cd ~

# Download the binary for your distribution
curl -O https://dl.google.com/go/go1.12.linux-amd64.tar.gz

# Verify that the downloaded binary is not corrupted. Check if the hash matches the one from downloads page
sha256sum go1.12.linux-amd64.tar.gz

# Extract the binary
tar xvf go1.12.linux-amd64.tar.gz

# Make root user the owner of Go workspace
sudo chown -R root:root ./go

# Move go directory to a standard location
sudo mv go /usr/local
```

For more details regarding `Go` installation check out

[Golang Installation](https://golang.org/doc/install)

For downloading `Go` binary for your platform checkout

[Golang Downloads](https://golang.org/dl/)

To check if you installed `Go` successfully type:

```bash
# Displays installed Go version
go version

# Displays all environment variables defined by Go
go env
```

Docker:

```bash
docker run -it golang:1.11
```

### Set $GOPATH variable
```bash
sudo nano ~/.profile

# Linux
export GOPATH=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/$GOBIN

# OSX
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/$GOBIN

# Save File

# Restart shell configuration
source ~/.profile
```

### Other Go tools

You can also install other `Go` tools which will help you have a more productive development process

- [Go Tools](https://github.com/golang/tools)
- [Go Lint](https://github.com/golang/lint)
- [Go Imports](https://godoc.org/golang.org/x/tools/cmd/goimports)

#### Go Tools

```bash
go get -u golang.org/x/tools/...
```

#### Go Lint

```bash
go get -u golang.org/x/lint/golint
```

#### Go Imports

```bash
go get -u golang.org/x/tools/cmd/goimports
```

#### Daily routine things

- Format your entire code recursively using
`go fmt ./...`

- Lint your entire code recursively using
`golint ./...` 

### For more info about Go check out

[Go Doc](https://golang.org/pkg)

[Go Packages](https://godoc.org/)

[Go Playground](https://play.golang.org/)

[Effective Go](https://golang.org/doc/effective_go.html)

[A Tour of Go](https://tour.golang.org/)


### For more info about Go commands & WORKSPACE check out

[Go commands explained](https://github.com/gophertuts/go-basics/blob/master/go-commands.md)

[Go WORKSPACE explained](https://github.com/gophertuts/go-basics/blob/master/go-workspace.md)

## FEEDBACK ⚗

[GopherTuts TypeForm](https://feedback.gophertuts.com)

## COMMUNITY 🙌

[GopherTuts Discord](https://discord.gg/4sgecdh)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="150px"/>