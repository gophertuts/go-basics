# Basic Types in Go

## Run code for `int` type

```bash
# cd into the project
cd $GOPATH/src/github.com/gophertuts/go-basics/basic-types/int

# 32bit
# prints 4  => 4B = 4 * 8b
env GOARCH=386 go run main.go

# 64bit
# prints 8 => 8B = 8 * 8b
env GOARCH=amd64 go run main.go

# if you know it's 64bit
go run main.go
```

## Resources

- [String implementation](https://github.com/golang/go/blob/master/src/runtime/string.go)
- [Strings - Go101](https://go101.org/article/string.html)
- [Unsafe package](https://golang.org/pkg/unsafe/)
- [Unsafe - Go101](https://go101.org/article/unsafe.html)
- [UintPtr vs Unsafe pointer](https://utcc.utoronto.ca/~cks/space/blog/programming/GoUintptrVsUnsafePointer)
