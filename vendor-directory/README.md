# `vendor` directory in Go

## Medium version

[Vendor directory in Go](https://medium.com/@gophertuts/packages-in-go-df5438123548)

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

## Resources

- [Go project structure](https://vsupalov.com/go-folder-structure/)
- [Go project layout](https://github.com/golang-standards/project-layout)
- [`go get` custom import path](https://jve.linuxwall.info/blog/index.php?post/2015/08/26/Hosting_Go_code_on_Github_with_custom_import_path)
- [Remote import paths](https://golang.org/cmd/go/#hdr-Remote_import_paths)

## FEEDBACK ⚗

[GopherTuts TypeForm](http://feedback.gophertuts.com)

Back to
[Go Basics](https://github.com/gophertuts/go-basics)

<img src="https://github.com/gophertuts/go-basics/raw/master/gophertuts.svg?sanitize=true" width="50px"/>
