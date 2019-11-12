# Custom package manager

## Medium version 📖

[Custom package manager in Go](https://medium.com/@gophertuts/packages-in-go-df5438123548)

#### How to test

```bash
# generate pkg-manager binary and places it inside $GOPATH/bin
go install

# export $GOPATH/bin and add it inside $PATH env variable
echo "export GOBIN=$GOPATH/bin\nexport PATH=$GOBIN:$PATH" >> .bashrc
echo "export GOBIN=$GOPATH/bin\nexport PATH=$GOBIN:$PATH" >> .zshrc

# create a simple go project
mkdir project-name
touch go-pkg.json
# or using
pkg-manager init project-name

# downloads all dependencies indicated in go-pkg.json
pkg-manager get ./...

# downloads the specified dependency and saves it inside in go-pkg.json
# using either the version tag in format: 'v1.2.3'
# or using the commit hash
pkg-manager get github.com/user/repo@version
pkg-manager get github.com/user/repo@GIT_COMMIT_HASH
```

#### Useful info

```bash
go help importpath
```

#### Improvements

- Add recursive dependency tree feature
- Add a lock file with detailed information about each dependency
- Add a caching mechanism
- Get rid of git checkout and only keep the code
- Add support for archives instead of VCS checkouts
- Add support for custom domains
- Add support for raw git URLs code fetching

## Resources 💎

- [Go 1.5 vendor experiment](https://go.googlesource.com/proposal/+/master/design/25719-go15vendor.md)
- [Vendor directories](https://golang.org/cmd/go/#hdr-Vendor_Directories)
- [VGO module](https://research.swtch.com/vgo-module)
- [Go project structure](https://vsupalov.com/go-folder-structure/)
- [Go project layout](https://github.com/golang-standards/project-layout)

## FEEDBACK ⚗

[GopherTuts TypeForm](http://feedback.gophertuts.com)

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>