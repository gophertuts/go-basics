# Go WORKSPACE & dependency management explained

Most of Go projects have a specific structure, which from project to project
is pretty much the same. The project normally lives in `$GOPATH` location,
which is the process developers do, right after installing Go binary.

Right after installing Go binary, it is recommended to correctly set
your `$GOPATH` variable, which most of the times, is set to
`$HOME/go`

The project layout represented below makes more sense once you start
using third party packages, or your own packages.

Traditionally third party packages which are hosted remotely
are usually installed using `go get` tool.

So for example installing `errors` package by `davecheney`
you simply run:
```bash
go get github.com/pkg/errors
```

This command will reach `github.com/pkg/errors` and place the source code
which will download inside `$GOPATH/src/github.com/pkg/errors`

If the package contains a main function it will also run
`go install` and create a binary for the downloaded package

#### Go project layout:
```
├── $GOPATH/
│   ├── src/
│   │   ├── github.com/
│   │   │   ├── organization/
│   │   │   │   ├── project1/
│   │   │   │   └── project2/
│   │   │   ├── user1/
│   │   │   │   ├── project1/
│   │   │   │   └── project2/
│   │   │   ├── pgk/
│   │   │   └── └── errors/
│   │   ├── gitlab.com/
│   │   │   ├── organization/
│   │   │   │   ├── project1/
│   │   │   │   └── project2/
│   │   │   ├── user1/
│   │   │   │   ├── project1/
│   │   │   │   └── project2/
│   │   │   ├── user2/
│   │   │   │   ├── project1/
│   │   │   └── └── project2/
│   │   ├── golang.org/
│   │   │   ├── x/
│   │   │   │   ├── crypto/
│   │   │   │   ├── lint/
│   │   │   └── └── sys/
│   ├── bin/
│   │   ├── goimports
│   │   ├── golint
│   │   ├── binary1
│   │   └── binary2
│   ├── pkg/
│   │   ├── darwin_amd64/
│   │   │   ├── github.com/
│   │   │   │   ├── uber-go/
│   │   │   │   │   └── atomic.a
│   │   │   │   ├── user1/
│   │   │   │   │   ├── package1/
│   │   │   └── └── └── package2/
│   │   ├── mod/
│   │   │   ├── github.com/
│   │   │   │   ├── uber-go/
│   │   │   │   │   └── zap@1.8.0
│   │   │   │   │   └── zap@1.9.1
│   │   │   └── └── └── ...
│   │   ├── dep/
│   │   │   ├── sources/
│   │   │   │   ├── https---github.com-pkg-errors/
│   │   │   │   │   ├── errors.go
│   │   │   │   │   ├── errors_test.go
│   │   │   │   │   └── ...
│   │   │   │   ├── https---github.com-user1-project1/
│   │   │   │   │   ├── file1.go
│   │   │   │   │   ├── file1_test.go
│   │   ├── ├── ├── └── ...
│   │   │   │   ├── https---github.com-user2-project1/
│   │   │   │   │   ├── file1.go
│   │   │   │   │   ├── file1_test.go
│   │   │   │   │   ├── file2.go
│   │   │   │   │   ├── file2_test.go
│   │   └── └── └── └── ...
│   └──
└──
```

### `$GOPATH`

`$GOPATH` is the location of Go workspace. In other words, the place
where all Go projects and third party packages will live.

It is recommended to set `$GOPATH` right after installing Go binary.
Also it is not recommended to have multiple `$GOPATH` variables
when switching projects.

By simply placing this inside your `.profile` or `.bashrc` or
`.zshrc` depending on which SHELL you use it should set `$GOPATH` correctly

```bash
export GOPATH=$HOME/go
```

Go `1.11` introduces modules, which allows users to create projects
and libraries outside `$GOPATH` location.

Most of the projects out there use `$GOPATH` and it's gonna stay there for a while.
Go modules are something new and definitely worth giving a try.

They seem to be a good feet for libraries, where less dependencies
are usually required, because Go modules come, with their own
configuration file and a verbosity layer

### `src`

Represents the directory where all the projects' source code and
third party packages will live. So when making a reference to a
third party package or one of your local packages, it will
reference `src` directory.

The source for a package with the import path `"X/Y/Z"` lives in the directory
`$GOPATH/src/X/Y/Z`

So importing `errors` package by `davecheney` will look something like:

```go
import (
	"os"
	"net/http"
	
	"github.com/pkg/errors"
	
	"github.com/username/project/package"
)
```

All of the references to `github.com/...` will be resolved down to
`src` directory where `github.com` lives.

Notice also that packages like: `os` and `net/http` do not require
this bloated prefix. It's because they are part of standard library.

Even local packages are referenced using this long path prefix,
hence `github.com/username/project/package`

Downloading third party packages using `go get` tool, will store packages
in the same `src` location using `git` to clone the repository.

### `bin`

Represents the location where all Go binaries are stored. Usually
when downloading a package using `go get` or when we manually run
`go install` inside the location where we have a main package, in
both cases these operations will result in creating a binary and
storing it inside `$GOPATH/bin` directory which is also represented
by `$GOBIN` environment variable.

The binary for a command whose source is in `$GOPATH/src/A/B` lives in
`$GOPATH/bin/B`

Also if `$GOBIN` variable is included in path, you can run any
compiled binary by simply typing their name.

So a good practice after installing Go binary is to include
`$GOBIN` inside `$PATH` variable so, every binary is easily accessible.

By simply placing this inside your `.profile` or `.bashrc` or
`.zshrc` depending on which SHELL you use it should set `$GOBIN` correctly.

```bash
export GOBIN=$GOPATH/bin
export PATH=$PATH:/$GOBIN
```

### `pkg`

Represents the location for package archives, where binaries
needed for a certain package import are stored. If a package is required by
`$GOPATH/src/github.com/username/project` then there is likely a chance,
that package will be stored as an archive file inside `$GOPATH/pkg`

The binary for a package with the import path `"X/Y/Z"` lives in
`$GOPATH/pkg/$GOOS_$GOARCH/X/Y/Z.a`

### `vendor` directory

The `vendor` directory inside `$GOPATH/host/username/project` is
recognized as a special directory by Go binary tool. And there are
plenty of tools such as `dep` which serve the purpose of package
management.

`$GOPATH` alone works nice, but when it comes to using different
versions of packages inside a project it fails to do so. 

Because by default `go get` tool will fetch the latest code from `master`
branch and that may be an issue is all of th sudden an issue/bug was
introduced in the latest version of a third party package, which
likely happens very often in the modern web development.

When Go encounters an import path like `"X/Y/Z"`, it will
try to resolve it in several ways. It will look for a `vendor`
directory first, if it finds any, it will resolve the package from
there, if not it will continue to look in `$GOPATH`.

For all this `vendor` directory magic to work, it is required
that your project is inside `$GOPATH`, otherwise this will
not work.

### `dep`

`Dep` is a modern Go package manager. For a long period of time
Go did not have any package managers. Meaning that the imported code
from third party packages was always latest version of it.

Any issue or bug introduced in the third party package would directly affect
the Go project without a way to rollback. Because that was a frequent
problem the community emerged with tools like`glide` and `dep` which are
package managers to address this issue.

To install `dep` check out the 
[Installation Guide](https://golang.github.io/dep/docs/installation.html)

Running `dep init` will get you started very quickly by generating
`Gopkg.toml` and `Gopkg.lock` files

Some daily `dep` commands:

```bash
# Installs all dependencies listed in "Gopkg.toml" and "Gopkg.lock"
dep ensure

# Ensures that "errors" & "bar" packages will be installed
dep ensure -add github.com/pkg/errors github.com/foo/bar

# Ensures that "bar" package will be updated
dep ensure -update github.com/foo/bar

# Reports if project and dependencies are out of sync
dep check
```
 
For more info on `dep` check out
[Dep Docs](https://golang.github.io/dep/docs/introduction.html)

### Go modules

Go modules are a new addition to in Go `1.11` where developers are
allowed to write Go code outside `$GOPATH`. This is not an article about
Go modules but the idea is simple, there is a configuration file `go.mod`
which tells Go how to resolve dependencies for a specific package.

Also a separate tool `go mod` is available for that purpose.

For more info check out

[Go Modules](https://github.com/golang/go/wiki/Modules#go-111-modules)

[Go Modules in 2019](https://blog.golang.org/modules2019)

### Notes:

Usually artifacts inside `bin` and `pkg` are generated and no
developer interaction is required. Basically all developers care about
is `$GOPATH/src` where the source code lives.

### For more info check out:

[How to write Go code](https://golang.org/doc/code.html)

[GOPATH](https://github.com/golang/go/wiki/GOPATH)

[Dep](https://golang.github.io/dep/docs/introduction.html)

[Go Modules](https://github.com/golang/go/wiki/Modules#go-111-modules)

Back to
[Go Basics](https://github.com/gophertuts/go-basics)
