# Custom package manager

#### How to test

```bash
# will generate pkg-manager binary located in $GOPATH/bin
go install

# add $GOPATH/bin inside $PATH

# create a simple go project

# downloads all dependencies indicated in go-pkg.json
pkg-manager get ./...

# downloads the specified dependency and saves it inside in go-pkg.json
# using either the version tag in format: 'v1.2.3'
# or using the commit hash
pkg-manager get github.com/user/repo@version
pkg-manager get github.com/user/repo@GIT_COMMIT_HASH
```
