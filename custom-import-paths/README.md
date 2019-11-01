# Custom **import** paths in Go

## Medium version

[Custom import paths in Go](https://medium.com/@gophertuts/packages-in-go-df5438123548)

### How to run:

1. Custom `go get` **import** path

```bash
# cd into the example
cd $GOPATH/src/github.com/gophertuts/go-basics/vendor-directory/go-get-custom-domain

# fetch the package from custom URL
go get -insecure 0.0.0.0/user/pkg

# run an HTTPS server
go run server.go
```

#### More info

Note: import path checking is disabled when using
Go modules

```bash
# display useful information about Go import paths
go help importpath

# display useful information about Go specific env variables
go help environment
```

## Resources

- [`go get` custom import path](https://jve.linuxwall.info/blog/index.php?post/2015/08/26/Hosting_Go_code_on_Github_with_custom_import_path)
- [Remote import paths](https://golang.org/cmd/go/#hdr-Remote_import_paths)
- [Custom import path checking](https://docs.google.com/document/d/1jVFkZTcYbNLaTxXD9OcGfn7vYv5hWtPx9--lTx1gPMs/edit)

## FEEDBACK ⚗

[GopherTuts TypeForm](http://feedback.gophertuts.com)

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>
