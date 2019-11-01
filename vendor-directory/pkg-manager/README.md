# Custom package manager

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

#### Improvements

- Add recursive dependency tree feature
- Add a lock file with detailed information about each dependency
- Add a caching mechanism
- Get rid of git checkout and only keep the code
- Add support for custom domains
- Add support for raw git URLs code fetching
