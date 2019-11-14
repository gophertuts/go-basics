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

- Add recursive dependency tree fetching
- Add a lock file with detailed dependency tree
- Add a caching mechanism and/or proxy mechanism
- Implement Go's download protocol
- Don't save git repositories, just the actual code
- Add support for archives instead of VCS checkouts
- Add the possibility to skip test files or non go files to save bandwidth
- Add support for custom domains & custom import paths
- Add support for multiple VCS systems

## Resources 💎

- [Go 1.5 vendor experiment](https://go.googlesource.com/proposal/+/master/design/25719-go15vendor.md)
- [Vendor directories](https://golang.org/cmd/go/#hdr-Vendor_Directories)
- [VGO module](https://research.swtch.com/vgo-module)
- [Go project structure](https://vsupalov.com/go-folder-structure/)
- [Go project layout](https://github.com/golang-standards/project-layout)
- [Dep package manager](https://github.com/golang/dep)
- [Project Athens - the download protocol](https://medium.com/@arschles/project-athens-the-download-protocol-2b346926a818)
- [Go module proxy protocol](https://golang.org/cmd/go/#hdr-Module_proxy_protocol)
- [Go module proxy - Fatih Arslan](https://arslan.io/2019/08/02/why-you-should-use-a-go-module-proxy/)
- [Go proxies](https://roberto.selbach.ca/go-proxies/)
- [Athens - GitHub project](https://medium.com/@arschles/project-athens-the-download-protocol-2b346926a818)
- [Go module sum](https://sum.golang.org/)

## FEEDBACK ⚗

[GopherTuts TypeForm](http://feedback.gophertuts.com)

## COMMUNITY 🙌

[GopherTuts Discord](https://discord.gg/4sgecdh)

---

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>