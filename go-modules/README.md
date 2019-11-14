# Go Modules

### `$GOPATH`

```bash
# display useful information about $GOPATH
go help gopath
```

### `$GOPRIVATE`

```bash
# Lists all available packages inside $GOPATH
# OR All packages inside current module
# if env variable GO111MODULE=on
go list all

# Lists all std lib packages and its sub packages
# Works the same as `go list all`
go list std

# Lists all cmd packages and its sub packages
# Works the same as `go list all`
go list cmd

# Lists all available local packages in CWD
go list ./...

# Lists all imported (local) packages in a JSON format
go list -json ./...

# Displays useful information about import paths
go help importpath
```

## Resources 💎

- [Migrating to Go modules](https://blog.golang.org/migrating-to-go-modules)
- [VGO module](https://research.swtch.com/vgo-module)
- [Module configuration (private modules)](https://golang.org/cmd/go/#hdr-Module_configuration_for_non_public_modules)

## FEEDBACK ⚗

[GopherTuts TypeForm](http://feedback.gophertuts.com)

## COMMUNITY 🙌

[GopherTuts Discord](https://discord.gg/4sgecdh)

---

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>
