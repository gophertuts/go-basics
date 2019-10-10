# `go get`

```bash
# Downloads all referenced dependencies (packages)
# by every package in current directory
go get ./...

# Downloads latest version (master) of 'pkg'
# with module aware feature on 
go get pkg

# Downloads v1.0.0 of 'pkg' package
go get pkg@v1.0.0

# Downloads 'pkg' package at git commit hash
go get pkg@af044c0995fe

# Locks the revision number for this package
# to prevent future breaking changes
go get github.com/lestrrat/go-jwx@b7d4802
cat go.mod

# Downloads or updates 'pkg' package
go get -u pkg

# For more info on go get
go help get
```

### Module aware feature

In order to turn this feature on you need
to export `GO111MODULE=on` env variable
inside your shell profile

# `go clean`

```bash
# cleans binaries generated with `go build`
# Note: it does not clean the binaries
# generated using `go build -o binary-name`
go clean

# For more info on go get
go help clean
```

### Best practices

- Use `go get ./...` when fetching all dependencies for
a project
- Use `go get pkg` for fetching the latest version
of a package
- Use `go get pkg@v1.0.0` for fetching a specific non
breaking version of the package
Note: `GO111MODULE=on` ENV var should be set to be able
to use module aware feature
